package sap_api_caller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-business-partner-creates-rmq-kube/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-business-partner-creates-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	outputQueues    []string
	outputter       RMQOutputter
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		outputQueues:    outputQueueTo,
		outputter:       outputter,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostBusinessPartner(
	address *requests.Address,
	bank *requests.Bank,
	general *requests.General,
	role *requests.Role,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Address":
			func() {
				c.Address(address)
				wg.Done()
			}()
		case "Bank":
			func() {
				c.Bank(bank)
				wg.Done()
			}()
		case "General":
			func() {
				c.General(general)
				wg.Done()
			}()
		case "Role":
			func() {
				c.Role(role)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Address(address *requests.Address) {
	addressData, err := c.callBusinessPartnerSrvAPIRequirementAddress("A_BusinessPartner", address)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": addressData, "function": "BusinessPartnerAddress"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(addressData)
}

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementAddress(api string, address *requests.Address) (*sap_api_output_formatter.Address, error) {
	body, err := json.Marshal(address)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Bank(bank *requests.Bank) {
	url := fmt.Sprintf("A_BusinessPartner", bank.BusinessPartner)
	outputDataBank, err := c.callBusinessPartnerSrvAPIRequirementBank(url, bank)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataBank, "function": "BusinessPartnerBank"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataBank)
}

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementBank(api string, bank *requests.Bank) (*sap_api_output_formatter.Bank, error) {
	body, err := json.Marshal(bank)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToBank(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) General(general *requests.General) {
	url := fmt.Sprintf("A_BusinessPartner", general.BusinessPartner)
	outputDataGeneral, err := c.callBusinessPartnerSrvAPIRequirementGeneral(url, general)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataGeneral, "function": "BusinessPartnerGeneral"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataGeneral)
}

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementGeneral(api string, general *requests.General) (*sap_api_output_formatter.General, error) {
	body, err := json.Marshal(general)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Role(role *requests.Role) {
	url := fmt.Sprintf("A_BusinessPartner", role.BusinessPartner)
	outputDataRole, err := c.callBusinessPartnerSrvAPIRequirementRole(url, role)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataRole, "function": "BusinessPartnerRole"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataRole)
}

func (c *SAPAPICaller) callBusinessPartnerSrvAPIRequirementRole(api string, role *requests.Role) (*sap_api_output_formatter.Role, error) {
	body, err := json.Marshal(role)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_BUSINESS_PARTNER_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToRole(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
