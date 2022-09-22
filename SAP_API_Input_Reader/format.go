package sap_api_input_reader

import (
	"sap-api-integrations-business-partner-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.BusinessPartner
	return &requests.General{
		BusinessPartner: data.BusinessPartner,
		//	Customer:                       data.Customer,
		//	Supplier:                       data.Supplier,
		AcademicTitle:           data.AcademicTitle,
		AuthorizationGroup:      data.AuthorizationGroup,
		BusinessPartnerCategory: data.BusinessPartnerCategory,
		//	BusinessPartnerFullName:        data.BusinessPartnerFullName,
		BusinessPartnerGrouping: data.BusinessPartnerGrouping,
		//	BusinessPartnerName:            data.BusinessPartnerName,
		CorrespondenceLanguage: data.CorrespondenceLanguage,
		//	CreationDate:                   data.CreationDate,
		//	CreationTime:                   data.CreationTime,
		FirstName:       data.FirstName,
		Industry:        data.Industry,
		IsFemale:        data.IsFemale,
		IsMale:          data.IsMale,
		IsNaturalPerson: data.IsNaturalPerson,
		IsSexUnknown:    data.IsSexUnknown,
		GenderCodeName:  data.GenderCodeName,
		Language:        data.Language,
		//	LastChangeDate:                 data.LastChangeDate,
		//	LastChangeTime:                 data.LastChangeTime,
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
		// BusinessPartnerDeathDate:       data.BusinessPartnerDeathDate,
		BusinessPartnerIsBlocked:  data.BusinessPartnerIsBlocked,
		BusinessPartnerType:       data.BusinessPartnerType,
		GroupBusinessPartnerName1: data.GroupBusinessPartnerName1,
		GroupBusinessPartnerName2: data.GroupBusinessPartnerName2,
		//	IndependentAddressID:           data.IndependentAddressID,
		MiddleName:                   data.MiddleName,
		NameCountry:                  data.NameCountry,
		PersonFullName:               data.PersonFullName,
		PersonNumber:                 data.PersonNumber,
		IsMarkedForArchiving:         data.IsMarkedForArchiving,
		BusinessPartnerIDByExtSystem: data.BusinessPartnerIDByExtSystem,
		TradingPartner:               data.TradingPartner,
	}
}

func (sdc *SDC) ConvertToRole() *requests.Role {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Role
	return &requests.Role{
		BusinessPartner:     dataBusinessPartner.BusinessPartner,
		BusinessPartnerRole: data.BusinessPartnerRole,
		ValidFrom:           data.ValidFrom,
		ValidTo:             data.ValidTo,
	}
}

func (sdc *SDC) ConvertToAddress() *requests.Address {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Address
	return &requests.Address{
		BusinessPartner:   dataBusinessPartner.BusinessPartner,
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
}

func (sdc *SDC) ConvertToBank() *requests.Bank {
	dataBusinessPartner := sdc.BusinessPartner
	data := sdc.BusinessPartner.Bank
	return &requests.Bank{
		BusinessPartner:          dataBusinessPartner.BusinessPartner,
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
}
