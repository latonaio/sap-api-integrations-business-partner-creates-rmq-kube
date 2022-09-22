package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-business-partner-creates-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) (*General, error) {
	pm := &responses.General{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	general := &General{
		BusinessPartner:               data.BusinessPartner,
		Customer:                      data.Customer,
		Supplier:                      data.Supplier,
		AcademicTitle:                 data.AcademicTitle,
		AuthorizationGroup:            data.AuthorizationGroup,
		BusinessPartnerCategory:       data.BusinessPartnerCategory,
		BusinessPartnerFullName:       data.BusinessPartnerFullName,
		BusinessPartnerGrouping:       data.BusinessPartnerGrouping,
		BusinessPartnerName:           data.BusinessPartnerName,
		CorrespondenceLanguage:        data.CorrespondenceLanguage,
		CreationDate:                  data.CreationDate,
		CreationTime:                  data.CreationTime,
		FirstName:                     data.FirstName,
		Industry:                      data.Industry,
		IsFemale:                      data.IsFemale,
		IsMale:                        data.IsMale,
		IsNaturalPerson:               data.IsNaturalPerson,
		IsSexUnknown:                  data.IsSexUnknown,
		GenderCodeName:                data.GenderCodeName,
		Language:                      data.Language,
		LastChangeDate:                data.LastChangeDate,
		LastChangeTime:                data.LastChangeTime,
		LastName:                      data.LastName,
		OrganizationBPName1:           data.OrganizationBPName1,
		OrganizationBPName2:           data.OrganizationBPName2,
		OrganizationBPName3:           data.OrganizationBPName3,
		OrganizationBPName4:           data.OrganizationBPName4,
		OrganizationFoundationDate:    data.OrganizationFoundationDate,
		OrganizationLiquidationDate:   data.OrganizationLiquidationDate,
		SearchTerm1:                   data.SearchTerm1,
		SearchTerm2:                   data.SearchTerm2,
		AdditionalLastName:            data.AdditionalLastName,
		BirthDate:                     data.BirthDate,
		BusinessPartnerBirthplaceName: data.BusinessPartnerBirthplaceName,
		BusinessPartnerDeathDate:      data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:      data.BusinessPartnerIsBlocked,
		BusinessPartnerType:           data.BusinessPartnerType,
		GroupBusinessPartnerName1:     data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2:     data.GroupBusinessPartnerName2,
		IndependentAddressID:          data.IndependentAddressID,
		MiddleName:                    data.MiddleName,
		NameCountry:                   data.NameCountry,
		PersonFullName:                data.PersonFullName,
		PersonNumber:                  data.PersonNumber,
		IsMarkedForArchiving:          data.IsMarkedForArchiving,
		BusinessPartnerIDByExtSystem:  data.BusinessPartnerIDByExtSystem,
		TradingPartner:                data.TradingPartner,
	}

	return general, nil
}

func ConvertToRole(raw []byte, l *logger.Logger) (*Role, error) {
	p := &responses.Role{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Role. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	role := &Role{
		BusinessPartner:     data.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
	}

	return role, nil
}

func ConvertToAddress(raw []byte, l *logger.Logger) (*Address, error) {
	p := &responses.Address{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Role. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	address := &Address{
		BusinessPartner:   data.BusinessPartner,
		AddressID:         data.AddressID,
		ValidityStartDate: data.ValidityStartDate,
		ValidityEndDate:   data.ValidityEndDate,
		Country:           data.Country,
		Region:            data.Region,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		PostalCode:        data.PostalCode,
		Language:          data.Language,
	}

	return address, nil
}

func ConvertToBank(raw []byte, l *logger.Logger) (*Bank, error) {
	p := &responses.Bank{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Bank. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	bank := &Bank{
		BusinessPartner:          data.BusinessPartner,
		BankIdentification:       data.BankIdentification,
		BankCountryKey:           data.BankCountryKey,
		BankName:                 data.BankName,
		BankNumber:               data.BankNumber,
		SWIFTCode:                data.SWIFTCode,
		BankControlKey:           data.BankControlKey,
		BankAccountHolderName:    data.BankAccountHolderName,
		BankAccountName:          data.BankAccountName,
		ValidityStartDate:        data.ValidityStartDate,
		ValidityEndDate:          data.ValidityEndDate,
		IBAN:                     data.IBAN,
		IBANValidityStartDate:    data.IBANValidityStartDate,
		BankAccount:              data.BankAccount,
		BankAccountReferenceText: data.BankAccountReferenceText,
		CollectionAuthInd:        data.CollectionAuthInd,
		CityName:                 data.CityName,
		AuthorizationGroup:       data.AuthorizationGroup,
	}

	return bank, nil
}
