package structuresEecASN1

import "local/structures/defStructs"

// TTtap310ZB
var TTtap310ZB = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTtap310ZB", 0x00ff0000, 0x00010000, 0, &Ttap310ZB},
	1: {"TTtap310ZB", 0x00ff0000, 0x00020000, 0, &Ttap310ZB},
}

var Ttap310ZB = map[int]defStructs.Champ{
	1: {"TransferBatch", -1, 0, &TrBatch},
	2: {"Notification", -1, 0, &Notif},
}

var TrBatch = map[int]defStructs.Champ{
	3:  {"CallEventDetailList", -1, 1, &CallEDListT2},
	4:  {"BatchControlInfo", -1, 1, &BatchContInfoT2},
	5:  {"AccountingInfo", -1, 1, &AccInfoT2},
	6:  {"NetworkInfo", -1, 1, &NetInfoT2},
	7:  {"VasInfoList", -1, 1, &VInfoListT2},
	8:  {"MessageDescriptionInforList", -1, 1, &MesDescInfoListT2},
	15: {"AuditControlInfo", -1, 1, &AuContInfoT2},
}

var Notif = map[int]defStructs.Champ{
	107: {"FileAvailableTimeStamp", -1, 1, &DateTimeLT3},
	108: {"FileCreationTimeStamp", -1, 1, &DateTimeLT3},
	109: {"FileSequenceNumber", -1, 2, nil},
	110: {"FileTypeIndicator", -1, 2, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT3},
	181: {"RapFileSequenceNumber", -1, 1, &RFileSeqNumbT3},
	182: {"Recipient", -1, 1, &RecT3},
	189: {"ReleaseVersionNumber", -1, 1, nil},
	196: {"Sender", -1, 1, &RecT3},
	201: {"SpecificationVersionNumber", -1, 1, nil},
	227: {"TransferCutOffTimeStamp", -1, 1, &DateTimeLT3},
}

var CallEDListT2 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, &MobOriginCallT3},
	10:  {"MobileTerminatedCall", -1, 1, &MobTermCallT3},
	11:  {"SupplServiceEvent", -1, 1, &SupplServEventT3},
	12:  {"ServiceCentreUsage", -1, 1, &ServCentreUsageT3},
	13:  {"ValueAddedService", -1, 1, &ValAddServiceT3},
	14:  {"GprsCall", -1, 1, &GCallT3},
	17:  {"ContentTransaction", -1, 1, &ContTransactionT3},
	297: {"LocationService", -1, 1, &LocServiceT3},
}

var BatchContInfoT2 = map[int]defStructs.Champ{
	107: {"FileAvailableTimeStamp", -1, 1, &DateTimeLT4},
	108: {"FileCreationTimeStamp", -1, 1, &DateTimeLT4},
	109: {"FileSequenceNumber", -1, 2, nil},
	110: {"FileTypeIndicator", -1, 2, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT3},
	181: {"RapFileSequenceNumber", -1, 1, &RFileSeqNumbT3},
	182: {"Recipient", -1, 2, &RecT3},
	189: {"ReleaseVersionNumber", -1, 1, nil},
	196: {"Sender", -1, 2, &RFileSeqNumbT3},
	201: {"SpecificationVersionNumber", -1, 1, nil},
	227: {"TransferCutOffTimeStamp", -1, 1, &DateTimeLT4},
}

var AccInfoT2 = map[int]defStructs.Champ{
	80:  {"CurrencyConversionList", -1, 1, &CurrConversionT3},
	95:  {"DiscountingList", -1, 1, &DiscListT3},
	135: {"LocalCurrency", -1, 2, nil},
	210: {"TapCurrency", -1, 2, nil},
	211: {"TaxationList", -1, 1, nil},
	244: {"TapDecimalPlaces", -1, 1, nil},
}

var NetInfoT2 = map[int]defStructs.Champ{
	157: {"NetworkType", -1, 1, nil},
	188: {"RecEntityInfoList", -1, 1, &RecEntInfListT3},
	234: {"UtcTimeOffsetInfoList", -1, 1, &UTOSILT3},
	249: {"CalledNumAnalysisList", -1, 1, &CalNumAnLystT3},
}

var VInfoListT2 = map[int]defStructs.Champ{
	238: {"VasInformation", -1, 1, &VasInfoT3},
}

var MesDescInfoListT2 = map[int]defStructs.Champ{
	143: {"MessageDescriptionInformation", -1, 1, &MessDescInfoT3},
}

var AuContInfoT2 = map[int]defStructs.Champ{
	43:  {"CallEventDetailsCount", -1, 1, nil},
	101: {"EarliestCallTimeStamp", -1, 1, &DateTimeLT3},
	133: {"LatestCallTimeStamp", -1, 1, &DateTimeLT3},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT3},
	224: {"TotalChargeValueList", -1, 1, &TotChValListT3},
	225: {"TotalDiscountValue", -1, 1, &TotalValueT3},
	226: {"TotalTaxValue", -1, 1, &TotalValueT3},
	353: {"TotalTaxRefund", -1, 1, &TotalValueT3},
	354: {"TotalDiscountRefund", -1, 1, &TotalValueT3},
	361: {"TotalAdvisedChargeValueList", -1, 1, &TotalAdvChValListT3},
}

var MobOriginCallT3 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT4},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT4},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT4},
	138: {"LocationInformation", -1, 1, &LocInfoT4},
	147: {"MoBasicCallInformation", -1, 1, &MoBasCallInfoT4},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT4},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT4},
}

var MobTermCallT3 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT4},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT4},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT4},
	138: {"LocationInformation", -1, 1, &LocInfoT4},
	153: {"MtBasicCallInformation", -1, 1, &MtBasCallInfoT4},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT4},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT4},
}

var SupplServEventT3 = map[int]defStructs.Champ{
	102: {"EquipmentInformation", -1, 1, &EquInfoT4},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	138: {"LocationInformation", -1, 1, &LocInfoT4},
	157: {"NetworkType", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT4},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT4},
	200: {"SimToolkitIndicator", -1, 2, nil},
	206: {"SupplServiceUsed", -1, 1, &SupplServUsedT4},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT4},
}

var ServCentreUsageT3 = map[int]defStructs.Champ{
	69:  {"ChargeInformation", -1, 1, &ChargInfoT4},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT4},
	184: {"RecEntityCode", -1, 1, &RecEntCodeT4},
	191: {"ScuBasicInformation", -1, 1, &ScuBasInfoT4},
	192: {"ScuChargeType", -1, 1, &ScuCharTypeT4},
	193: {"ScuTimeStamps", -1, 1, &ScuTimeStampsT4},
	195: {"ServingNetwork", -1, 2, nil},
}

var ValAddServiceT3 = map[int]defStructs.Champ{
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT4},
	198: {"ServingBid", -1, 1, &ServBidT4},
	226: {"TotalTaxValue", -1, 1, &TotalValueT4},
	236: {"ValueAddedServiceUsed", -1, 1, &ValAddServUsedT4},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT4},
}

var GCallT3 = map[int]defStructs.Champ{
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT4},
	102: {"EquipmentInformation", -1, 1, &EquInfoT4},
	114: {"GprsBasiCallInformation", -1, 1, &GprsBasCallInfoT4},
	117: {"GprsLocationInformation", -1, 1, &GprsLocInfoT4},
	121: {"GprsServiceUsed", -1, 1, &GprsServUsedT4},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	229: {"TypeOfControllingNode", -1, 1, nil},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT4},
}

var ContTransactionT3 = map[int]defStructs.Champ{
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	285: {"ContentServiceUsedList", -1, 1, &ContServUsedListT4},
	304: {"ContentTransactionBasicInfo", -1, 1, &ContTransBasicInfoT4},
	324: {"ChargedPartyInformation", -1, 1, &CharPartyInfoT4},
	335: {"ServingPartiesInformation", -1, 1, &ServPartInfoT4},
}

var LocServiceT3 = map[int]defStructs.Champ{
	45:  {"CallReference", -1, 1, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT4},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT4},
	184: {"RecEntityCode", -1, 1, &RecEntCodeT4},
	298: {"TrackingCustomerInformation", -1, 1, &TrackCustomInfoT4},
	367: {"TrackedCustomerInformation", -1, 1, &TrackedCustomInfoT4},
	373: {"LCSSPInformation", -1, 1, &LCSSPInfoT4},
	382: {"LocationServiceUsage", -1, 1, &LocServUsT4},
}

var TimeStampT3 = map[int]defStructs.Champ{
	84: {"DateTimeLong", -1, 1, &DateTimeLT4},
}

var OpSpecInfoT3 = map[int]defStructs.Champ{
	163: {"OperatorSpecInformation", -1, 2, nil},
}

var RFileSeqNumbT3 = map[int]defStructs.Champ{
	109: {"FileSequenceNumber", -1, 2, nil},
}

var RecT3 = map[int]defStructs.Champ{
	169: {"PlmnId", -1, 2, nil},
}

var DiscListT3 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	92: {"DiscountRate", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT4},
}

var RecEntInfListT3 = map[int]defStructs.Champ{
	183: {"RecEntityInformation", -1, 1, &RecEntInfoT4},
}

var UTOSILT3 = map[int]defStructs.Champ{
	233: {"UtcTimeOffsetInfo", -1, 1, &UtcTimeOffInfoT4},
}

var CalNumAnLystT3 = map[int]defStructs.Champ{
	246: {"CalledNumAnalysis", -1, 1, &CallNumAnT4},
}

var VasInfoT3 = map[int]defStructs.Champ{
	237: {"VasCode", -1, 1, &RecEntCodeT4},
	239: {"VasDescription", -1, 2, nil},
	240: {"VasShortDescription", -1, 2, nil},
}

var MessDescInfoT3 = map[int]defStructs.Champ{
	141: {"MessageDescriptionCode", -1, 1, &RecEntCodeT4},
	142: {"MessageDescription", -1, 1, nil},
}

var TotChValListT3 = map[int]defStructs.Champ{
	222: {"TotalChargeValue", -1, 1, &TotChargValT4},
}

var TotalValueT3 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var TotalAdvChValListT3 = map[int]defStructs.Champ{
	360: {"TotalAdvisedChargeValue", -1, 1, &TotAdCharValT4},
}

var CurrConversionT3 = map[int]defStructs.Champ{
	106: {"CurrencyConversion", -1, 1, &CurrConversionT4},
}

var DateTimeLT3 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	231: {"UtcTimeOffset", -1, 2, nil},
}

var CurrConversionT4 = map[int]defStructs.Champ{
	104: {"ExchangeRate", -1, 1, nil},
	105: {"Exchange", -1, 1, &RecEntCodeT5},
	159: {"NumberOfDecimalPlaces", -1, 1, nil},
}

var TotalValueT4 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var BasServUsedListT4 = map[int]defStructs.Champ{
	39: {"BasicServiceUsed", -1, 1, &BasServUsedT5},
}

var CamServUsedT4 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	49:  {"CamelCallReference", -1, 1, &CamCallRefT5},
	52:  {"CamelInitiatedCFIndicator", -1, 1, nil},
	54:  {"CamelModificationList", -1, 1, &CamModifListT5},
	55:  {"CamelServiceKey", -1, 1, nil},
	56:  {"CamelServiceLevel", -1, 1, nil},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT5},
	79:  {"CseInformation", -1, 1, nil},
	87:  {"DefaultCallHandlingIndicator", -1, 1, nil},
	116: {"GprsDestination", -1, 1, &GprsDestT5},
	399: {"CamelServerAddress", -1, 1, &CamServAddT5},
}

var EquInfoT4 = map[int]defStructs.Champ{
	103: {"Esn", -1, 1, nil},
	128: {"Imei", -1, 0, nil},
	148: {"MobileStationClassMark", -1, 1, nil},
}

var LocInfoT4 = map[int]defStructs.Champ{
	113: {"GeographicLocation", -1, 1, &GeoLocT5},
	123: {"HomeLocationInformation", -1, 1, &HomeLocInfoT5},
	156: {"NetworkLocation", -1, 1, &NetworkLocT5},
}

var MoBasCallInfoT4 = map[int]defStructs.Champ{
	44:  {"CallEventStartTimeStamp", -1, 1, &DateT5},
	58:  {"CauseForTerm", -1, 1, nil},
	89:  {"Destination", -1, 1, &DestinT5},
	90:  {"DestinationNetwork", -1, 1, &DestNetworkT5},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT5},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT5},
	200: {"SimToolkitIndicator", -1, 2, nil},
	223: {"TotalCallEventDuration", -1, 7, nil},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT5},
}

var OpSpecInfoT4 = map[int]defStructs.Champ{
	163: {"OperatorSpecInformation", -1, 2, nil},
}

var SupplServUsedListT4 = map[int]defStructs.Champ{
	206: {"SupplServiceUsed", -1, 1, &SupplServUsedT5},
}

var ValAddServUsedListT4 = map[int]defStructs.Champ{
	236: {"ValueAddedServiceUsed", -1, 1, &ValAddServUsedT5},
}

var MtBasCallInfoT4 = map[int]defStructs.Champ{
	41:  {"CallOriginator", -1, 1, &CallOriginT5},
	44:  {"CallEventStartTimeStamp", -1, 1, &DateT5},
	58:  {"CauseForTerm", -1, 1, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	164: {"OriginatingNetwork", -1, 1, &DestNetworkT5},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT5},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT5},
	200: {"SimToolkitIndicator", -1, 2, nil},
	223: {"TotalCallEventDuration", -1, 7, nil},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT5},
}

var RapFileSeqNumberT4 = map[int]defStructs.Champ{
	109: {"FileSequenceNumber", -1, 2, nil},
}

var SimChargSubT4 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 1, nil},
}

var SupplServUsedT4 = map[int]defStructs.Champ{
	37:  {"BasicServiceCodeList", -1, 1, &BasicServCodeListT5},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT5},
	74:  {"ChargingTimeStamp", -1, 1, &DateT5},
	204: {"SsParameters", -1, 2, nil},
	208: {"SupplServiceActionCod", -1, 1, nil},
	209: {"SupplServiceCode", -1, 2, nil},
	219: {"ThirdPartyInformation", -1, 1, &CallOriginT5},
}

var MinChargSubT4 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var ChargInfoT4 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetailList", -1, 1, &ChargDetListT5},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT5},
	105: {"Exchange", -1, 1, &RecEntCodeT5},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT5},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT5},
}

var RecEntCodeT4 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var ScuBasInfoT4 = map[int]defStructs.Champ{
	67:  {"ChargedPartyStatus", -1, 1, nil},
	75:  {"ClirIndicator", -1, 1, nil},
	90:  {"DestinationNetwork", -1, 1, &DestNetworkT5},
	158: {"NonChargedParty", -1, 1, &NonChargPartyT5},
	164: {"OriginatingNetwork", -1, 1, &DestNetworkT5},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT5},
	286: {"GSMChargeableSubscriber", -1, 1, &GSMChargSubT5},
}

var ScuCharTypeT4 = map[int]defStructs.Champ{
	98:  {"DistanceChargeBandCode", -1, 2, nil},
	141: {"MessageDescriptionCode", -1, 1, &RecEntCodeT5},
	144: {"MessageStatus", -1, 1, nil},
	145: {"MessageNumber", -1, 1, nil},
	170: {"PriorityCode", -1, 1, nil},
}

var ScuTimeStampsT4 = map[int]defStructs.Champ{
	73: {"ChargingPoint", -1, 2, nil},
	76: {"CompletionTimeStamp", -1, 1, &DateT5},
	88: {"DepositTimeStamp", -1, 1, &DateT5},
}

var ServBidT4 = map[int]defStructs.Champ{
	242: {"Bid", -1, 2, nil},
}

var ValAddServUsedT4 = map[int]defStructs.Champ{
	69:  {"ChargeInformation", -1, 1, &ChargInfoT5},
	74:  {"ChargingTimeStamp", -1, 1, &DateT5},
	237: {"VasCode", -1, 1, &RecEntCodeT5},
}

var GprsBasCallInfoT4 = map[int]defStructs.Champ{
	44:  {"CallEventStartTimeStamp", -1, 1, &DateT5},
	58:  {"CauseForTerm", -1, 1, nil},
	72:  {"CharginfId", -1, 0, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	115: {"GprsChargeableSubscriber", -1, 1, &GprsChargSubT5},
	116: {"GprsDestination", -1, 1, &GprsDestT5},
	157: {"NetworkType", -1, 1, nil},
	166: {"PartialTypeIndicator", -1, 2, nil},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT5},
	223: {"TotalCallEventDuration", -1, 7, nil},
	245: {"NetworkInitPDPContent", -1, 1, nil},
}

var GprsLocInfoT4 = map[int]defStructs.Champ{
	113: {"GeographicLocation", -1, 1, &GeoLocT5},
	118: {"GprsNetworkLocation", -1, 1, &GprsNetworkLocT5},
	123: {"HomeLocationInformation", -1, 1, &HomeLocInfoT5},
}

var GprsServUsedT4 = map[int]defStructs.Champ{
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT5},
	74:  {"ChargingTimeStamp", -1, 1, &DateT5},
	120: {"GprsServiceUsageList", -1, 1, &GprsServUsListT5},
}

var ContServUsedListT4 = map[int]defStructs.Champ{
	352: {"ContentServiceUsed", -1, 1, &ContServUsedT5},
}

var ContTransBasicInfoT4 = map[int]defStructs.Champ{
	112: {"FraudMonitorIndicator", -1, 2, nil},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT5},
	223: {"TotalCallEventDuration", -1, 7, nil},
	300: {"OrderPlaceTimeStamp", -1, 1, &DateT5},
	301: {"RequestedDeliveryTimeStamp", -1, 1, &DateT5},
	302: {"ActualDeliveryTimeStamp", -1, 1, &DateT5},
	303: {"TransactionStatus", -1, 1, nil},
}

var CharPartyInfoT4 = map[int]defStructs.Champ{
	310: {"ChargedPartyIdList", -1, 1, &ChargPartListT5},
	314: {"ChargedPartyHomeIdList", -1, 1, &ChargPartHomeIdListT5},
	321: {"ChargedPartyLocationList", -1, 1, &ChargPartLocListT5},
	323: {"ChargedPartyEquip", -1, 1, &ChargPartEquipT5},
}

var ServPartInfoT4 = map[int]defStructs.Champ{
	155: {"NetworkId", -1, 2, nil},
	328: {"ContentProviderIdList", -1, 1, &ContProvIdListT5},
	330: {"InternetServiceProviderIdList", -1, 1, nil},
	334: {"ContentProviderName", -1, 2, nil},
}

var TrackCustomInfoT4 = map[int]defStructs.Champ{
	299: {"TrackingCustomerIdList", -1, 1, &TrackCustomIdListT5},
	365: {"TrackingCustomerHomeIdList", -1, 1, &TrackCustomHomeListT5},
	368: {"TrackingCustomerLocList", -1, 1, &TrackCustomLocListT5},
	371: {"TrackingCustomerEquipment", -1, 1, &TrackCustomEquipT5},
}

var TrackedCustomInfoT4 = map[int]defStructs.Champ{
	370: {"TrackedCustomerIdList", -1, 1, &TrackedCustomIdListT5},
	376: {"TrackedCustomerHomeIdList", -1, 1, &TrackedCustomHoListT5},
	379: {"TrackedCustomerLocList", -1, 1, &TrackedCustomLocListT5},
	381: {"TrackingCustomerEquipment", -1, 1, &TrackCustomEquipT5},
}

var LCSSPInfoT4 = map[int]defStructs.Champ{
	333: {"NetworkList", -1, 1, &NetListT5},
	374: {"LCSSPIdentificationList", -1, 1, &LCSSPIdListT5},
	378: {"ISPList", -1, 1, &ISPListT5},
}

var LocServUsT4 = map[int]defStructs.Champ{
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT5},
	74:  {"ChargingTimeStamp", -1, 1, &DateT5},
	383: {"LCSQosRequested", -1, 1, &LCSQosReqT5},
	390: {"LCSQosDelivered", -1, 1, &LCSQosDelivT5},
}

var DateTimeLT4 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	231: {"UtcTimeOffset", -1, 2, nil},
}

var RecEntInfoT4 = map[int]defStructs.Champ{
	131: {"IPTextV4Address", -1, 2, nil},
	132: {"IPTextV6Address", -1, 2, nil},
	151: {"MscId", -1, 2, nil},
	152: {"Msisdn", -1, 0, nil},
	184: {"RecEntityCode", -1, 1, &RecEntCodeT5},
	186: {"RecEntityType", -1, 1, nil},
	296: {"GmlcAddress", -1, 2, nil},
}

var UtcTimeOffInfoT4 = map[int]defStructs.Champ{
	231: {"UtcTimeOffset", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT5},
}

var CallNumAnT4 = map[int]defStructs.Champ{
	78:  {"CountryCodeList", -1, 1, &CountCodeListT5},
	127: {"IacList", -1, 1, &IacListT5},
	247: {"CalledNumAnalysisCode", -1, 1, &RecEntCodeT5},
}

var TotChargValT4 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	241: {"AbsoluteAmount", -1, 1, nil},
	355: {"TotalChargeRefund", -1, 1, &TotalValueT5},
}

var TotAdCharValT4 = map[int]defStructs.Champ{
	348: {"AdvisedChargeCurrency", -1, 2, nil},
	356: {"TotalAdvisedCharge", -1, 1, &TotalValueT5},
	357: {"TotalAdvisedChargeRefund", -1, 1, &TotalValueT5},
	358: {"TotalCommission", -1, 1, &TotalValueT5},
	359: {"TotalCommissionRefund", -1, 1, &TotalValueT5},
}

var TaxInfoListT5 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, &MobOriginCallT6},
	10:  {"MobileTerminatedCall", -1, 1, &MobTermCallT6},
	14:  {"GprsCall", -1, 1, &GCallT6},
	15:  {"AuditControlInfo", -1, 1, &AuContInfoT6},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT6},
	213: {"TaxInformation", -1, 1, &TaxInfoT6},
}

var ChargDetListT5 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT6},
}

var BasServUsedT5 = map[int]defStructs.Champ{
	36:  {"BasicService", -1, 1, &BasiCServT6},
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT6},
	74:  {"ChargingTimeStamp", -1, 1, &DateT6},
	82:  {"DataVolumeReference", -1, 1, nil},
	124: {"HSCSDInformation", -1, 1, &HSCSDinfoT6},
}

var CamCallRefT5 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	50: {"CamelCallReferenceNumber", -1, 2, nil},
}

var CamModifListT5 = map[int]defStructs.Champ{
	53: {"CamelModification", -1, 1, nil},
}

var ChargInfoT5 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetailList", -1, 1, &ChargDetListT6},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT6},
	105: {"Exchange", -1, 1, &RecEntCodeT6},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT6},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT6},
}

var TotalValueT5 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var GprsDestT5 = map[int]defStructs.Champ{
	190: {"RemotePdpAddressList", -1, 1, &RemPdpAddListT6},
	261: {"AccessPointNameNI", -1, 2, nil},
	262: {"AccessPointNameOI", -1, 2, nil},
}

var CamServAddT5 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
}

var GeoLocT5 = map[int]defStructs.Champ{
	134: {"Latitude", -1, 1, nil},
	137: {"LocationDescription", -1, 2, nil},
	139: {"Longitude", -1, 1, nil},
	195: {"ServingNetwork", -1, 2, nil},
	198: {"ServingBid", -1, 1, &ServBidT6},
}

var HomeLocInfoT5 = map[int]defStructs.Champ{
	122: {"HomeBid", -1, 1, &ServBidT6},
	137: {"LocationDescription", -1, 2, nil},
}

var NetworkLocT5 = map[int]defStructs.Champ{
	45:  {"CallReference", -1, 1, nil},
	59:  {"CellId", -1, 0, nil},
	136: {"LocationArea", -1, 0, nil},
	184: {"RecEntityCode", -1, 1, &RecEntCodeT6},
}

var DateT5 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT6},
}

var DestinT5 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	42:  {"CalledPlace", -1, 2, nil},
	46:  {"CalledRegion", -1, 2, nil},
	247: {"CalledNumAnalysisCode", -1, 1, &RecEntCodeT6},
	279: {"DialledDigits", -1, 2, nil},
}

var DestNetworkT5 = map[int]defStructs.Champ{
	155: {"NetworkId", -1, 2, nil},
}

var RapFileSeqNumberT5 = map[int]defStructs.Champ{
	109: {"FileSequenceNumber", -1, 2, nil},
}

var SimChargSubT5 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var MinChargSubT5 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var SupplServUsedT5 = map[int]defStructs.Champ{
	37:  {"BasicServiceCodeList", -1, 1, &BasicServCodeListT6},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT6},
	74:  {"ChargingTimeStamp", -1, 1, &DateT6},
	204: {"SsParameters", -1, 2, nil},
	208: {"SupplServiceActionCod", -1, 1, nil},
	209: {"SupplServiceCode", -1, 2, nil},
	219: {"ThirdPartyInformation", -1, 1, &CallOriginT6},
}

var ValAddServUsedT5 = map[int]defStructs.Champ{
	69:  {"ChargeInformation", -1, 1, &ChargInfoT6},
	74:  {"ChargingTimeStamp", -1, 1, &DateT6},
	237: {"VasCode", -1, 1, &RecEntCodeT6},
}

var CallOriginT5 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	75: {"ClirIndicator", -1, 1, nil},
}

var BasicServCodeListT5 = map[int]defStructs.Champ{
	40:  {"BearerServiceCode", -1, 2, nil},
	218: {"TeleServiceCode", -1, 2, nil},
}

var RecEntCodeT5 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT5 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, &TotalValueT6},
	398: {"TaxableAmount", -1, 1, &TotalValueT6},
}

var CallTypeGrT5 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var NonChargPartyT5 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	247: {"CalledNumAnalysisCode", -1, 1, &RecEntCodeT6},
}

var GSMChargSubT5 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var GprsChargSubT5 = map[int]defStructs.Champ{
	167: {"PdpAddress", -1, 2, &PdpAddT6},
	168: {"PdpType", -1, 1, nil},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT6},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT6},
	263: {"ChargingCharacteristics", -1, 1, nil},
}

var GprsNetworkLocT5 = map[int]defStructs.Champ{
	59:  {"CellId", -1, 0, nil},
	136: {"LocationArea", -1, 0, nil},
	185: {"RecEntityCodeList", -1, 1, &RecEntCodeListT6},
}

var GprsServUsListT5 = map[int]defStructs.Champ{
	119: {"GprsServiceUsage", -1, 1, &GprsServUsT6},
}

var ContServUsedT5 = map[int]defStructs.Champ{
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT6},
	250: {"DataVolumeIncoming", -1, 0, &DataVolT6},
	251: {"DataVolumeOutgoing", -1, 0, &DataVolT6},
	281: {"ObjectType", -1, 1, nil},
	336: {"ContentTransactionCode", -1, 1, nil},
	337: {"ContentTransactionType", -1, 1, nil},
	338: {"TransactionDescriptionSupp", -1, 1, nil},
	339: {"TransactionDetailDescription", -1, 2, nil},
	340: {"TransactionShortDescription", -1, 2, nil},
	341: {"TransactionIdentifier", -1, 2, nil},
	342: {"TransactionAuthCode", -1, 2, nil},
	343: {"TotalDataVolume", -1, 1, &DataVolT6},
	344: {"ChargeRefundIndicator", -1, 1, nil},
	345: {"ContentChargingPoint", -1, 1, nil},
	351: {"AdvisedChargeInformation", -1, 1, &AdvChargInfoT6},
}

var ChargPartListT5 = map[int]defStructs.Champ{
	287: {"ChargedPartyIdentifier", -1, 2, nil},
}

var ChargPartHomeIdListT5 = map[int]defStructs.Champ{
	313: {"ChargedPartyHomeIdentification", -1, 1, &ChargPartHomeIdT6},
}

var ChargPartLocListT5 = map[int]defStructs.Champ{
	320: {"ChargedPartyLocation", -1, 1, &ChargPartLocT6},
}

var ChargPartEquipT5 = map[int]defStructs.Champ{
	322: {"EquipmentIdType", -1, 1, nil},
	290: {"EquipmentId", -1, 2, nil},
}

var ChargInfoListT5 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT6},
}

var DiscInfoT5 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT6},
}

var ContProvIdListT5 = map[int]defStructs.Champ{
	327: {"ContentProvider", -1, 1, &ContProvT6},
}

var TrackCustomIdListT5 = map[int]defStructs.Champ{
	362: {"TrackingCustomerIdentification", -1, 1, &TrackCustomIdT6},
}

var TrackCustomHomeListT5 = map[int]defStructs.Champ{
	366: {"TrackingCustomerHomeId", -1, 1, &ChargPartHomeIdT6},
}

var TrackCustomLocListT5 = map[int]defStructs.Champ{
	369: {"TrackingCustomerLocation", -1, 1, &ChargPartLocT6},
}

var TrackCustomEquipT5 = map[int]defStructs.Champ{
	290: {"EquipmentId", -1, 2, nil},
	322: {"EquipmentIdType", -1, 1, nil},
}

var TrackedCustomIdListT5 = map[int]defStructs.Champ{
	372: {"TrackedCustomerIdentification", -1, 1, &TrackCustomIdT6},
}

var TrackedCustomHoListT5 = map[int]defStructs.Champ{
	377: {"TrackedCustomerHomeId", -1, 1, &ChargPartHomeIdT6},
}

var TrackedCustomLocListT5 = map[int]defStructs.Champ{
	380: {"TrackedCustomerLocation", -1, 1, &ChargPartLocT6},
}

var NetListT5 = map[int]defStructs.Champ{
	332: {"Network", -1, 1, &NetworkT6},
}

var LCSSPIdListT5 = map[int]defStructs.Champ{
	375: {"LCSSPIdentificationList", -1, 1, &ContProvT6},
}

var ISPListT5 = map[int]defStructs.Champ{
	329: {"InternetServiceProvider", -1, 1, &IntServProvT6},
}

var LCSQosReqT5 = map[int]defStructs.Champ{
	384: {"LCSRequestTimeStamp", -1, 1, &DateT6},
	385: {"HorizontalAccuracyRequested", -1, 1, nil},
	386: {"VerticalAccuracyRequested", -1, 1, nil},
	387: {"PesponseTimeCategory", -1, 1, nil},
	388: {"TrackingPeriod", -1, 1, nil},
	389: {"TrackingFrequency", -1, 1, nil},
}

var LCSQosDelivT5 = map[int]defStructs.Champ{
	388: {"TrackingPeriod", -1, 1, nil},
	389: {"TrackingFrequency", -1, 1, nil},
	392: {"HorizontalAccuracyDelivered", -1, 1, nil},
	393: {"VerticalAccuracyDelivered", -1, 1, nil},
	394: {"ResponseTime", -1, 1, nil},
}

var CountCodeListT5 = map[int]defStructs.Champ{
	77: {"CountryCode", -1, 2, nil},
}

var IacListT5 = map[int]defStructs.Champ{
	126: {"Iac", -1, 2, nil},
}

var MobOriginCallT6 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT7},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT7},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT7},
	138: {"LocationInformation", -1, 1, &LocInfoT7},
	147: {"MoBasicCallInformation", -1, 1, &MoBasCallInfoT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT7},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT7},
}

var MobTermCallT6 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT7},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT7},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT7},
	138: {"LocationInformation", -1, 1, &LocInfoT7},
	153: {"MtBasicCallInformation", -1, 1, &MtBasCallInfoT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT7},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT7},
}

var GCallT6 = map[int]defStructs.Champ{
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT7},
	102: {"EquipmentInformation", -1, 1, &EquInfoT7},
	114: {"GprsBasiCallInformation", -1, 1, &GprsBasCallInfoT7},
	117: {"GprsLocationInformation", -1, 1, &GprsLocInfoT7},
	121: {"GprsServiceUsed", -1, 1, &GprsServUsedT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	229: {"TypeOfControllingNode", -1, 1, nil},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT7},
}

var AuContInfoT6 = map[int]defStructs.Champ{
	43:  {"CallEventDetailsCount", -1, 1, nil},
	101: {"EarliestCallTimeStamp", -1, 1, &DateTimeLT7},
	133: {"LatestCallTimeStamp", -1, 1, &DateTimeLT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	224: {"TotalChargeValueList", -1, 1, &TotChValListT7},
	225: {"TotalDiscountValue", -1, 1, &TotalValueT7},
	226: {"TotalTaxValue", -1, 1, &TotalValueT7},
	353: {"TotalTaxRefund", -1, 1, &TotalValueT7},
	354: {"TotalDiscountRefund", -1, 1, &TotalValueT7},
	361: {"TotalAdvisedChargeValueList", -1, 1, &TotalAdvChValListT7},
}

var OpSpecInfoT6 = map[int]defStructs.Champ{
	163: {"OperatorSpecInformation", -1, 2, nil},
}

var TaxInfoListT6 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, &MobOriginCallT7},
	10:  {"MobileTerminatedCall", -1, 1, &MobTermCallT7},
	14:  {"GprsCall", -1, 1, &GCallT7},
	15:  {"AuditControlInfo", -1, 1, &AuContInfoT7},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	213: {"TaxInformation", -1, 1, &TaxInfoT7},
}

var ChargDetListT6 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT7},
}

var ChargDetT6 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT7},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT7},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var TotalValueT6 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateT6 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT7},
}

var ChargInfoT6 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetailList", -1, 1, &ChargDetListT7},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT7},
	105: {"Exchange", -1, 1, &RecEntCodeT7},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT7},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT7},
}

var ChargInfoListT6 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT7},
}

var DiscInfoT6 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT7},
}

var RecEntCodeT6 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT6 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, &TotalValueT7},
	398: {"TaxableAmount", -1, 1, &TotalValueT7},
}

var CallTypeGrT6 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var BasiCServT6 = map[int]defStructs.Champ{
	34:  {"AiurRequested", -1, 1, nil},
	40:  {"BearerServiceCode", -1, 2, nil},
	111: {"Fnur", -1, 1, nil},
	179: {"RadioChannetRequested", -1, 1, nil},
	180: {"RadioChannelUsed", -1, 1, nil},
	202: {"SpeedVersionRequested", -1, 1, nil},
	203: {"SpeedVersionUsed", -1, 1, nil},
	218: {"TeleServiceCode", -1, 2, nil},
	228: {"TransparencyIndicator", -1, 1, nil},
	280: {"UserProtocolIndicator", -1, 1, nil},
}

var HSCSDinfoT6 = map[int]defStructs.Champ{
	35:  {"BasicHSCSDParameters", -1, 1, &BasHSCDParamT7},
	140: {"HSCSDParameterModificationList", -1, 1, &HSCSDParamModListT7},
}

var RemPdpAddListT6 = map[int]defStructs.Champ{
	165: {"PacketDataProtocolAddress", -1, 2, nil},
}

var ServBidT6 = map[int]defStructs.Champ{
	242: {"Bid", -1, 2, nil},
}

var DateTimeT6 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT7},
}

var BasicServCodeListT6 = map[int]defStructs.Champ{
	40:  {"BearerServiceCode", -1, 2, nil},
	218: {"TeleServiceCode", -1, 2, nil},
}

var CallOriginT6 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	75: {"ClirIndicator", -1, 1, nil},
}

var PdpAddT6 = map[int]defStructs.Champ{
	165: {"PacketDataProtocolAddress", -1, 2, nil},
}

var SimChargSubT6 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var MinChargSubT6 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var RecEntCodeListT6 = map[int]defStructs.Champ{
	184: {"RecEntityCode", -1, 1, &RecEntCodeT7},
}

var GprsServUsT6 = map[int]defStructs.Champ{
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT7},
	250: {"DataVolumeIncoming", -1, 0, &DataVolT7},
	251: {"DataVolumeOutgoing", -1, 0, &DataVolT7},
	270: {"UMTSQoSRequested", -1, 1, &UMTSQoST7},
	271: {"UMTSQoSUsed", -1, 1, &UMTSQoST7},
	272: {"GSMQoSRequested", -1, 1, &GSMQosT7},
	273: {"GSMQoSUsed", -1, 1, &GSMQosT7},
}

var DataVolT6 = map[int]defStructs.Champ{
	81: {"DataVolume", -1, 1, nil},
}

var AdvChargInfoT6 = map[int]defStructs.Champ{
	346: {"PaidIndicator", -1, 1, nil},
	347: {"PaymentMethod", -1, 1, nil},
	348: {"AdvisedChargeCurrency", -1, 2, nil},
	349: {"AdvisedCharge", -1, 1, &ChargeT7},
	350: {"Commission", -1, 1, &ChargeT7},
}

var ChargPartHomeIdT6 = map[int]defStructs.Champ{
	288: {"HomeIdentifier", -1, 2, nil},
	311: {"HomeIdType", -1, 1, nil},
}

var ChargPartLocT6 = map[int]defStructs.Champ{
	289: {"LocationIdentifier", -1, 2, nil},
	315: {"LocationIdType", -1, 1, nil},
}

var ContProvT6 = map[int]defStructs.Champ{
	291: {"ContentProviderIdType", -1, 1, nil},
	292: {"ContentProviderIdentifier", -1, 2, nil},
}

var TrackCustomIdT6 = map[int]defStructs.Champ{
	363: {"CustomerIdType", -1, 1, nil},
	364: {"CustomerIdentifier", -1, 2, nil},
}

var NetworkT6 = map[int]defStructs.Champ{
	295: {"NetworkIdentifier", -1, 2, nil},
	331: {"NetworkIdType", -1, 1, nil},
}

var IntServProvT6 = map[int]defStructs.Champ{
	293: {"IspIdType", -1, 1, nil},
	294: {"IspIdentifier", -1, 2, nil},
}

var MobOriginCallT7 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT7},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT7},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT7},
	138: {"LocationInformation", -1, 1, &LocInfoT7},
	147: {"MoBasicCallInformation", -1, 1, &MoBasCallInfoT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT7},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT7},
}

var MobTermCallT7 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT7},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT7},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT7},
	138: {"LocationInformation", -1, 1, &LocInfoT7},
	153: {"MtBasicCallInformation", -1, 1, &MtBasCallInfoT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT7},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT7},
}

var GCallT7 = map[int]defStructs.Champ{
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT7},
	102: {"EquipmentInformation", -1, 1, &EquInfoT7},
	114: {"GprsBasiCallInformation", -1, 1, &GprsBasCallInfoT7},
	117: {"GprsLocationInformation", -1, 1, &GprsLocInfoT7},
	121: {"GprsServiceUsed", -1, 1, &GprsServUsedT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	229: {"TypeOfControllingNode", -1, 1, nil},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT7},
}

var AuContInfoT7 = map[int]defStructs.Champ{
	43:  {"CallEventDetailsCount", -1, 1, nil},
	101: {"EarliestCallTimeStamp", -1, 1, &DateTimeLT7},
	133: {"LatestCallTimeStamp", -1, 1, &DateTimeLT7},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT7},
	224: {"TotalChargeValueList", -1, 1, &TotChValListT7},
	225: {"TotalDiscountValue", -1, 1, &TotalValueT7},
	226: {"TotalTaxValue", -1, 1, &TotalValueT7},
	353: {"TotalTaxRefund", -1, 1, &TotalValueT7},
	354: {"TotalDiscountRefund", -1, 1, &TotalValueT7},
	361: {"TotalAdvisedChargeValueList", -1, 1, &TotalAdvChValListT7},
}

var OpSpecInfoT7 = map[int]defStructs.Champ{
	163: {"OperatorSpecInformation", -1, 2, nil},
}

var MtBasCallInfoT7 = map[int]defStructs.Champ{
	41:  {"CallOriginator", -1, 1, &CallOriginT9},
	44:  {"CallEventStartTimeStamp", -1, 1, &DateTimeT8},
	58:  {"CauseForTerm", -1, 1, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	164: {"OriginatingNetwork", -1, 1, &DestNetworkT8},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT8},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT8},
	200: {"SimToolkitIndicator", -1, 2, nil},
	223: {"TotalCallEventDuration", -1, 7, nil},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT8},
}

var BasServUsedListT7 = map[int]defStructs.Champ{
	39: {"BasicServiceUsed", -1, 1, &BasServUsedT8},
}

var GprsBasCallInfoT7 = map[int]defStructs.Champ{
	44:  {"CallEventStartTimeStamp", -1, 1, &DateTimeT8},
	58:  {"CauseForTerm", -1, 1, nil},
	72:  {"CharginfId", -1, 0, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	115: {"GprsChargeableSubscriber", -1, 1, &GprsChargSubT8},
	116: {"GprsDestination", -1, 1, &GprsDestT8},
	157: {"NetworkType", -1, 1, nil},
	166: {"PartialTypeIndicator", -1, 2, nil},
	181: {"RapFileSequenceNumber", -1, 17, &RapFileSeqNumberT8},
	223: {"TotalCallEventDuration", -1, 7, nil},
	245: {"NetworkInitPDPContent", -1, 1, nil},
}

var GprsLocInfoT7 = map[int]defStructs.Champ{
	113: {"GeographicLocation", -1, 1, &GeoLocT8},
	118: {"GprsNetworkLocation", -1, 1, &GprsNetworkLocT8},
	123: {"HomeLocationInformation", -1, 1, &HomeLocInfoT8},
}

var GprsServUsedT7 = map[int]defStructs.Champ{
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT8},
	74:  {"ChargingTimeStamp", -1, 1, &DateT8},
	120: {"GprsServiceUsageList", -1, 1, &GprsServUsListT8},
}

var CamServUsedT7 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	49:  {"CamelCallReference", -1, 1, &CamCallRefT8},
	52:  {"CamelInitiatedCFIndicator", -1, 1, nil},
	54:  {"CamelModificationList", -1, 1, &CamModifListT8},
	55:  {"CamelServiceKey", -1, 1, nil},
	56:  {"CamelServiceLevel", -1, 1, nil},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT8},
	79:  {"CseInformation", -1, 1, nil},
	87:  {"DefaultCallHandlingIndicator", -1, 1, nil},
	116: {"GprsDestination", -1, 1, &GprsDestT8},
	399: {"CamelServerAddress", -1, 1, &CamServAddT8},
}

var EquInfoT7 = map[int]defStructs.Champ{
	103: {"Esn", -1, 1, nil},
	128: {"Imei", -1, 0, nil},
	148: {"MobileStationClassMark", -1, 1, nil},
}

var LocInfoT7 = map[int]defStructs.Champ{
	113: {"GeographicLocation", -1, 1, &GeoLocT8},
	123: {"HomeLocationInformation", -1, 1, &HomeLocInfoT8},
	156: {"NetworkLocation", -1, 1, &NetworkLocT8},
}

var MoBasCallInfoT7 = map[int]defStructs.Champ{
	44:  {"CallEventStartTimeStamp", -1, 1, &DateTimeT8},
	58:  {"CauseForTerm", -1, 1, nil},
	89:  {"Destination", -1, 1, &DestinT8},
	90:  {"DestinationNetwork", -1, 1, &DestNetworkT8},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT8},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT8},
	200: {"SimToolkitIndicator", -1, 2, nil},
	223: {"TotalCallEventDuration", -1, 7, nil},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT8},
}

var SupplServUsedListT7 = map[int]defStructs.Champ{
	206: {"SupplServiceUsed", -1, 1, &SupplServUsedT8},
}

var ValAddServUsedListT7 = map[int]defStructs.Champ{
	236: {"ValueAddedServiceUsed", -1, 1, &ValAddServUsedT8},
}

var TotalAdvChValListT7 = map[int]defStructs.Champ{
	360: {"TotalAdvisedChargeValue", -1, 1, &TotAdCharValT8},
}

var TotChValListT7 = map[int]defStructs.Champ{
	222: {"TotalChargeValue", -1, 1, &TotChargValT8},
}

var DateTimeLT7 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	231: {"UtcTimeOffset", -1, 2, nil},
}

var TaxInfoListT7 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, nil},
	10:  {"MobileTerminatedCall", -1, 1, nil},
	14:  {"GprsCall", -1, 1, nil},
	15:  {"AuditControlInfo", -1, 1, nil},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, nil},
	213: {"TaxInformation", -1, 1, &TaxInfoT8},
}

var ChargDetListT7 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT8},
}

var ChargDetT7 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT8},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT8},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var ChargInfoT7 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetailList", -1, 1, &ChargDetListT8},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT8},
	105: {"Exchange", -1, 1, &RecEntCodeT8},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT8},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT8},
}

var ChargInfoListT7 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT8},
}

var DiscInfoT7 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT8},
}

var RecEntCodeT7 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT7 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, &TotalValueT8},
	398: {"TaxableAmount", -1, 1, &TotalValueT8},
}

var CallTypeGrT7 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var TotalValueT7 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateTimeT7 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT8},
}

var BasHSCDParamT7 = map[int]defStructs.Champ{
	60:  {"ChannelCoding", -1, 1, nil},
	61:  {"ChallelAcceptableList", -1, 1, &ChannelCodT8},
	161: {"NumberOfChannels", -1, 1, nil},
	252: {"NumberOfChannelUsed", -1, 1, nil},
}

var HSCSDParamModListT7 = map[int]defStructs.Champ{
	125: {"HSCSDParameterModification", -1, 1, &HSCSDParamModT8},
}

var DateT7 = map[int]defStructs.Champ{
	83: {"DateTime", -1, 1, &DateTimeT8},
}

var DataVolT7 = map[int]defStructs.Champ{
	81: {"DataVolume", -1, 1, nil},
}

var UMTSQoST7 = map[int]defStructs.Champ{
	264: {"QoSMaxBitRateUplink", -1, 2, nil},
	268: {"QoSTrafficClass", -1, 1, nil},
	274: {"QoSMaxBitRateDownlink", -1, 1, nil},
	275: {"QoSAlloRetenPriority", -1, 1, nil},
	283: {"QoSGuaranteedBitRateDownlink", -1, 2, nil},
	284: {"QoSGuaranteedBitRateUplink", -1, 2, nil},
}

var GSMQosT7 = map[int]defStructs.Champ{
	171: {"QoSDelay", -1, 1, nil},
	173: {"QoSMeanThrougtput", -1, 1, nil},
	174: {"QoSPeakThrougtput", -1, 1, nil},
	175: {"QoSPrecedence", -1, 1, nil},
	176: {"QoSreliability", -1, 1, nil},
}

var ChargeT7 = map[int]defStructs.Champ{
	62: {"Charge", -1, 1, &TotalValueT8},
}

var BasServUsedT8 = map[int]defStructs.Champ{
	36:  {"BasicService", -1, 1, &BasiCServT9},
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT9},
	74:  {"ChargingTimeStamp", -1, 1, &DateT9},
	82:  {"DataVolumeReference", -1, 1, nil},
	124: {"HSCSDInformation", -1, 1, &HSCSDinfoT9},
}

var GprsChargSubT8 = map[int]defStructs.Champ{
	167: {"PdpAddress", -1, 2, &PdpAddT9},
	168: {"PdpType", -1, 1, nil},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT9},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT9},
	263: {"ChargingCharacteristics", -1, 1, nil},
}

var GprsNetworkLocT8 = map[int]defStructs.Champ{
	59:  {"CellId", -1, 0, nil},
	136: {"LocationArea", -1, 0, nil},
	185: {"RecEntityCodeList", -1, 1, &RecEntCodeListT9},
}

var GprsServUsListT8 = map[int]defStructs.Champ{
	119: {"GprsServiceUsage", -1, 1, &GprsServUsT9},
}

var ValAddServUsedT8 = map[int]defStructs.Champ{
	69:  {"ChargeInformation", -1, 1, &ChargInfoT9},
	74:  {"ChargingTimeStamp", -1, 1, &DateT9},
	237: {"VasCode", -1, 1, &RecEntCodeT9},
}

var SupplServUsedT8 = map[int]defStructs.Champ{
	37:  {"BasicServiceCodeList", -1, 1, &BasicServCodeListT9},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT9},
	74:  {"ChargingTimeStamp", -1, 1, &DateT9},
	204: {"SsParameters", -1, 2, nil},
	208: {"SupplServiceActionCod", -1, 1, nil},
	209: {"SupplServiceCode", -1, 2, nil},
	219: {"ThirdPartyInformation", -1, 1, &CallOriginT9},
}

var MinChargSubT8 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var SimChargSubT8 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var RapFileSeqNumberT8 = map[int]defStructs.Champ{
	109: {"FileSequenceNumber", -1, 2, nil},
}

var DestNetworkT8 = map[int]defStructs.Champ{
	155: {"NetworkId", -1, 2, nil},
}

var DestinT8 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	42:  {"CalledPlace", -1, 2, nil},
	46:  {"CalledRegion", -1, 2, nil},
	247: {"CalledNumAnalysisCode", -1, 1, &RecEntCodeT9},
	279: {"DialledDigits", -1, 2, nil},
}

var DateT8 = map[int]defStructs.Champ{
	83: {"DateTime", -1, 1, &DateTimeT9},
}

var NetworkLocT8 = map[int]defStructs.Champ{
	45:  {"CallReference", -1, 1, nil},
	59:  {"CellId", -1, 0, nil},
	136: {"LocationArea", -1, 0, nil},
	184: {"RecEntityCode", -1, 1, &RecEntCodeT9},
}

var HomeLocInfoT8 = map[int]defStructs.Champ{
	122: {"HomeBid", -1, 1, &ServBidT9},
	137: {"LocationDescription", -1, 2, nil},
}

var GeoLocT8 = map[int]defStructs.Champ{
	134: {"Latitude", -1, 1, nil},
	137: {"LocationDescription", -1, 2, nil},
	139: {"Longitude", -1, 1, nil},
	195: {"ServingNetwork", -1, 2, nil},
	198: {"ServingBid", -1, 1, &ServBidT9},
}

var CamServAddT8 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
}

var GprsDestT8 = map[int]defStructs.Champ{
	190: {"RemotePdpAddressList", -1, 1, &RemPdpAddListT9},
	261: {"AccessPointNameNI", -1, 2, nil},
	262: {"AccessPointNameOI", -1, 2, nil},
}

var CamModifListT8 = map[int]defStructs.Champ{
	53: {"CamelModification", -1, 1, nil},
}

var CamCallRefT8 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	50: {"CamelCallReferenceNumber", -1, 2, nil},
}

var TotAdCharValT8 = map[int]defStructs.Champ{
	348: {"AdvisedChargeCurrency", -1, 2, nil},
	356: {"TotalAdvisedCharge", -1, 1, &TotalValueT9},
	357: {"TotalAdvisedChargeRefund", -1, 1, &TotalValueT9},
	358: {"TotalCommission", -1, 1, &TotalValueT9},
	359: {"TotalCommissionRefund", -1, 1, &TotalValueT9},
}

var TotChargValT8 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	241: {"AbsoluteAmount", -1, 1, nil},
	355: {"TotalChargeRefund", -1, 1, &TotalValueT9},
}

var TaxInfoListT8 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, &MobOriginCallT10},
	10:  {"MobileTerminatedCall", -1, 1, &MobTermCallT10},
	14:  {"GprsCall", -1, 1, &GCallT10},
	15:  {"AuditControlInfo", -1, 1, &AuContInfoT10},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT10},
	213: {"TaxInformation", -1, 1, &TaxInfoT9},
}

var ChargDetListT8 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT9},
}

var ChargDetT8 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT9},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT9},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var ChargInfoT8 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetailList", -1, 1, &ChargDetListT9},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT9},
	105: {"Exchange", -1, 1, &RecEntCodeT9},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT9},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT9},
}

var ChargInfoListT8 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT9},
}
var DiscInfoT8 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT9},
}

var RecEntCodeT8 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT8 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, &TotalValueT9},
	398: {"TaxableAmount", -1, 1, &TotalValueT9},
}

var CallTypeGrT8 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var TotalValueT8 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateTimeT8 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT9},
}

var ChannelCodT8 = map[int]defStructs.Champ{
	60: {"ChannelCoding", -1, 1, nil},
}

var HSCSDParamModT8 = map[int]defStructs.Champ{
	34:  {"AiurRequested", -1, 1, nil},
	60:  {"ChannelCoding", -1, 1, nil},
	130: {"InitiatingParty", -1, 1, nil},
	150: {"ModificationTimeStamp", -1, 1, &DateT9},
	161: {"NumberOfChannels", -1, 1, nil},
	252: {"NumberOfChannelUsed", -1, 1, nil},
}

var HSCSDinfoT9 = map[int]defStructs.Champ{
	35:  {"BasicHSCSDParameters", -1, 1, &BasHSCDParamT10},
	140: {"HSCSDParameterModificationList", -1, 1, &HSCSDParamModListT10},
}

var BasiCServT9 = map[int]defStructs.Champ{
	34:  {"AiurRequested", -1, 1, nil},
	40:  {"BearerServiceCode", -1, 2, nil},
	111: {"Fnur", -1, 1, nil},
	179: {"RadioChannetRequested", -1, 1, nil},
	180: {"RadioChannelUsed", -1, 1, nil},
	202: {"SpeedVersionRequested", -1, 1, nil},
	203: {"SpeedVersionUsed", -1, 1, nil},
	218: {"TeleServiceCode", -1, 2, nil},
	228: {"TransparencyIndicator", -1, 1, nil},
	280: {"UserProtocolIndicator", -1, 1, nil},
}

var MinChargSubT9 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var SimChargSubT9 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var PdpAddT9 = map[int]defStructs.Champ{
	165: {"PacketDataProtocolAddress", -1, 2, nil},
}

var RecEntCodeListT9 = map[int]defStructs.Champ{
	184: {"RecEntityCode", -1, 1, &RecEntCodeT10},
}

var GprsServUsT9 = map[int]defStructs.Champ{
	74:  {"ChargingTimeStamp", -1, 1, &DateT10},
	250: {"DataVolumeIncoming", -1, 0, &DataVolT10},
	251: {"DataVolumeOutgoing", -1, 0, &DataVolT10},
	270: {"UMTSQoSRequested", -1, 1, &UMTSQoST10},
	271: {"UMTSQoSUsed", -1, 1, &UMTSQoST10},
	272: {"GSMQoSRequested", -1, 1, &GSMQosT10},
	273: {"GSMQoSUsed", -1, 1, &GSMQosT10},
}

var BasicServCodeListT9 = map[int]defStructs.Champ{
	40:  {"BearerServiceCode", -1, 2, nil},
	218: {"TeleServiceCode", -1, 2, nil},
}

var CallOriginT9 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	75: {"ClirIndicator", -1, 1, nil},
}

var ServBidT9 = map[int]defStructs.Champ{
	242: {"Bid", -1, 2, nil},
}

var RemPdpAddListT9 = map[int]defStructs.Champ{
	165: {"PacketDataProtocolAddress", -1, 2, nil},
}

var TaxInfoListT9 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, &MobOriginCallT10},
	10:  {"MobileTerminatedCall", -1, 1, &MobTermCallT10},
	14:  {"GprsCall", -1, 1, &GCallT10},
	15:  {"AuditControlInfo", -1, 1, &AuContInfoT10},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT10},
	213: {"TaxInformation", -1, 1, &TaxInfoT10},
}

var ChargDetListT9 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT10},
}

var ChargDetT9 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT10},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT10},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var ChargInfoT9 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetailList", -1, 1, &ChargDetListT10},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT10},
	105: {"Exchange", -1, 1, &RecEntCodeT10},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT10},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT10},
}

var ChargInfoListT9 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT10},
}

var DiscInfoT9 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT10},
}

var RecEntCodeT9 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT9 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, &TotalValueT10},
	398: {"TaxableAmount", -1, 1, &TotalValueT10},
}

var CallTypeGrT9 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var TotalValueT9 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateT9 = map[int]defStructs.Champ{
	83: {"DateTime", -1, 1, &DateTimeT10},
}

var DateTimeT9 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT10},
}

var MobOriginCallT10 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT11},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT11},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT11},
	138: {"LocationInformation", -1, 1, &LocInfoT11},
	147: {"MoBasicCallInformation", -1, 1, &MoBasCallInfoT11},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT11},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT11},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT11},
}

var MobTermCallT10 = map[int]defStructs.Champ{
	38:  {"BasicServiceUsedList", -1, 1, &BasServUsedListT11},
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT11},
	99:  {"DualBearerServiceCode", -1, 1, nil},
	100: {"DualTeleServiceCode", -1, 1, nil},
	102: {"EquipmentInformation", -1, 1, &EquInfoT11},
	138: {"LocationInformation", -1, 1, &LocInfoT11},
	153: {"MtBasicCallInformation", -1, 1, &MtBasCallInfoT11},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT11},
	207: {"SupplServiceUsedList", -1, 1, &SupplServUsedListT11},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT11},
}

var GCallT10 = map[int]defStructs.Champ{
	57:  {"CamelServiceUsed", -1, 1, &CamServUsedT11},
	102: {"EquipmentInformation", -1, 1, &EquInfoT11},
	114: {"GprsBasiCallInformation", -1, 1, &GprsBasCallInfoT11},
	117: {"GprsLocationInformation", -1, 1, &GprsLocInfoT11},
	121: {"GprsServiceUsed", -1, 1, &GprsServUsedT11},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT11},
	229: {"TypeOfControllingNode", -1, 1, nil},
	235: {"ValueAddedServiceUsedList", -1, 1, &ValAddServUsedListT11},
}

var AuContInfoT10 = map[int]defStructs.Champ{
	43:  {"CallEventDetailsCount", -1, 1, nil},
	101: {"EarliestCallTimeStamp", -1, 1, &DateTimeLT11},
	133: {"LatestCallTimeStamp", -1, 1, &DateTimeLT11},
	162: {"OperatorSpecInfoList", -1, 1, &OpSpecInfoT11},
	224: {"TotalChargeValueList", -1, 1, &TotChValListT11},
	225: {"TotalDiscountValue", -1, 1, &TotalValueT11},
	226: {"TotalTaxValue", -1, 1, &TotalValueT11},
	353: {"TotalTaxRefund", -1, 1, &TotalValueT11},
	354: {"TotalDiscountRefund", -1, 1, &TotalValueT11},
	361: {"TotalAdvisedChargeValueList", -1, 1, &TotalAdvChValListT11},
}

var OpSpecInfoT10 = map[int]defStructs.Champ{
	163: {"OperatorSpecInformation", -1, 2, nil},
}

var HSCSDParamModListT10 = map[int]defStructs.Champ{
	125: {"HSCSDParameterModification", -1, 1, &HSCSDParamModT11},
}

var BasHSCDParamT10 = map[int]defStructs.Champ{
	60:  {"ChannelCoding", -1, 1, nil},
	61:  {"ChallelAcceptableList", -1, 1, &ChannelCodT11},
	161: {"NumberOfChannels", -1, 1, nil},
	252: {"NumberOfChannelUsed", -1, 1, nil},
}

var DateT10 = map[int]defStructs.Champ{
	83: {"DateTime", -1, 1, &DateTimeT11},
}

var DataVolT10 = map[int]defStructs.Champ{
	81: {"DataVolume", -1, 1, nil},
}

var UMTSQoST10 = map[int]defStructs.Champ{
	264: {"QoSMaxBitRateUplink", -1, 2, nil},
	268: {"QoSTrafficClass", -1, 1, nil},
	274: {"QoSMaxBitRateDownlink", -1, 1, nil},
	275: {"QoSAlloRetenPriority", -1, 1, nil},
	283: {"QoSGuaranteedBitRateDownlink", -1, 2, nil},
	284: {"QoSGuaranteedBitRateUplink", -1, 2, nil},
}

var GSMQosT10 = map[int]defStructs.Champ{
	171: {"QoSDelay", -1, 1, nil},
	173: {"QoSMeanThrougtput", -1, 1, nil},
	174: {"QoSPeakThrougtput", -1, 1, nil},
	175: {"QoSPrecedence", -1, 1, nil},
	176: {"QoSreliability", -1, 1, nil},
}

var TaxInfoListT10 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, nil},
	10:  {"MobileTerminatedCall", -1, 1, nil},
	14:  {"GprsCall", -1, 1, nil},
	15:  {"AuditControlInfo", -1, 1, nil},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, nil},
	213: {"TaxInformation", -1, 1, &TaxInfoT11},
}

var ChargDetListT10 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT11},
}

var ChargDetT10 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT11},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT11},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var ChargInfoT10 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetList", -1, 1, &ChargDetListT11},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT11},
	105: {"Exchange", -1, 1, &RecEntCodeT11},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT11},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT11},
}

var ChargInfoListT10 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT11},
}

var DiscInfoT10 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, &TotalValueT11},
}

var RecEntCodeT10 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT10 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, &TotalValueT11},
	398: {"TaxableAmount", -1, 1, &TotalValueT11},
}

var CallTypeGrT10 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var TotalValueT10 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateTimeT10 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT11},
}

var GprsBasCallInfoT11 = map[int]defStructs.Champ{
	44:  {"CallEventStartTimeStamp", -1, 1, &DateT12},
	58:  {"CauseForTerm", -1, 1, nil},
	72:  {"CharginfId", -1, 0, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	115: {"GprsChargeableSubscriber", -1, 1, &GprsChargSubT12},
	116: {"GprsDestination", -1, 1, &GprsDestT12},
	157: {"NetworkType", -1, 1, nil},
	166: {"PartialTypeIndicator", -1, 2, nil},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT12},
	223: {"TotalCallEventDuration", -1, 7, nil},
	245: {"NetworkInitPDPContent", -1, 1, nil},
}

var GprsLocInfoT11 = map[int]defStructs.Champ{
	113: {"GeographicLocation", -1, 1, &GeoLocT12},
	118: {"GprsNetworkLocation", -1, 1, &GprsNetworkLocT12},
	123: {"HomeLocationInformation", -1, 1, &HomeLocInfoT12},
}

var GprsServUsedT11 = map[int]defStructs.Champ{
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT13},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT13},
	120: {"GprsServiceUsageList", -1, 1, &GprsServUsListT12},
}

var TotalAdvChValListT11 = map[int]defStructs.Champ{
	360: {"TotalAdvisedChargeValue", -1, 1, &TotAdCharValT12},
}

var TotChValListT11 = map[int]defStructs.Champ{
	222: {"TotalChargeValue", -1, 1, &TotChargValT12},
}

var DateTimeLT11 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	231: {"UtcTimeOffset", -1, 2, nil},
}

var ValAddServUsedListT11 = map[int]defStructs.Champ{
	236: {"ValueAddedServiceUsed", -1, 1, &ValAddServUsedT12},
}

var SupplServUsedListT11 = map[int]defStructs.Champ{
	206: {"SupplServiceUsed", -1, 1, &SupplServUsedT12},
}

var OpSpecInfoT11 = map[int]defStructs.Champ{
	163: {"OperatorSpecInformation", -1, 2, nil},
}

var MtBasCallInfoT11 = map[int]defStructs.Champ{
	41:  {"CallOriginator", -1, 1, &CallOriginT12},
	44:  {"CallEventStartTimeStamp", -1, 1, &DateT12},
	58:  {"CauseForTerm", -1, 1, nil},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	164: {"OriginatingNetwork", -1, 1, &DestNetworkT12},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT12},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT12},
	200: {"SimToolkitIndicator", -1, 2, nil},
	223: {"TotalCallEventDuration", -1, 7, nil},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT12},
}

var MoBasCallInfoT11 = map[int]defStructs.Champ{
	44:  {"CallEventStartTimeStamp", -1, 1, &DateT12},
	58:  {"CauseForTerm", -1, 1, nil},
	89:  {"Destination", -1, 1, &DestinT12},
	90:  {"DestinationNetwork", -1, 1, &DestNetworkT12},
	112: {"FraudMonitorIndicator", -1, 2, nil},
	157: {"NetworkType", -1, 1, nil},
	181: {"RapFileSequenceNumber", -1, 1, &RapFileSeqNumberT12},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT12},
	200: {"SimToolkitIndicator", -1, 2, nil},
	223: {"TotalCallEventDuration", -1, 7, nil},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT12},
}

var EquInfoT11 = map[int]defStructs.Champ{
	103: {"Esn", -1, 1, nil},
	128: {"Imei", -1, 0, nil},
	148: {"MobileStationClassMark", -1, 1, nil},
}

var LocInfoT11 = map[int]defStructs.Champ{
	113: {"GeographicLocation", -1, 1, &GeoLocT12},
	123: {"HomeLocationInformation", -1, 1, &HomeLocInfoT12},
	156: {"NetworkLocation", -1, 1, &NetworkLocT12},
}

var CamServUsedT11 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	49:  {"CamelCallReference", -1, 1, &CamCallRefT12},
	52:  {"CamelInitiatedCFIndicator", -1, 1, nil},
	54:  {"CamelModificationList", -1, 1, &CamModifListT12},
	55:  {"CamelServiceKey", -1, 1, nil},
	56:  {"CamelServiceLevel", -1, 1, nil},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT14},
	79:  {"CseInformation", -1, 1, nil},
	87:  {"DefaultCallHandlingIndicator", -1, 1, nil},
	116: {"GprsDestination", -1, 1, &GprsDestT12},
	399: {"CamelServerAddress", -1, 1, &CamServAddT12},
}

var BasServUsedListT11 = map[int]defStructs.Champ{
	39: {"BasicServiceUsed", -1, 1, &BasServUsedT12},
}

var HSCSDParamModT11 = map[int]defStructs.Champ{
	34:  {"AiurRequested", -1, 1, nil},
	60:  {"ChannelCoding", -1, 1, nil},
	130: {"InitiatingParty", -1, 1, nil},
	150: {"ModificationTimeStamp", -1, 1, &DateT12},
	161: {"NumberOfChannels", -1, 1, nil},
	252: {"NumberOfChannelUsed", -1, 1, nil},
}

var TaxInfoListT11 = map[int]defStructs.Champ{
	213: {"TaxInformation", -1, 1, &TaxInfoT12},
}

var ChargDetListT11 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, nil},
}

var ChargDetT11 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT12},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, nil},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var ChannelCodT11 = map[int]defStructs.Champ{
	60: {"ChannelCoding", -1, 1, nil},
}

var DateTimeT11 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT12},
}

var ChargInfoT11 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	70:  {"ChargeInformationList", -1, 1, nil},
	96:  {"DiscountInformation", -1, 1, nil},
	105: {"Exchange", -1, 1, nil},
	214: {"TaxInformationList", -1, 1, nil},
	258: {"CallTypeGroup", -1, 1, nil},
}

var ChargInfoListT11 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, nil},
}

var DiscInfoT11 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, nil},
}

var RecEntCodeT11 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT11 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, nil},
	398: {"TaxableAmount", -1, 1, nil},
}

var CallTypeGrT11 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var TotalValueT11 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var GprsServUsListT12 = map[int]defStructs.Champ{
	119: {"GprsServiceUsage", -1, 1, &GprsServUsT13},
}

var GprsNetworkLocT12 = map[int]defStructs.Champ{
	59:  {"CellId", -1, 0, nil},
	136: {"LocationArea", -1, 0, nil},
	185: {"RecEntityCodeList", -1, 1, &RecEntCodeListT13},
}

var GprsChargSubT12 = map[int]defStructs.Champ{
	167: {"PdpAddress", -1, 2, &PdpAddT13},
	168: {"PdpType", -1, 1, nil},
	199: {"SimChargeableSubscriber", -1, 1, &SimChargSubT13},
	254: {"MinChargeableSubscriber", -1, 1, &MinChargSubT13},
	263: {"ChargingCharacteristics", -1, 1, nil},
}

var TotAdCharValT12 = map[int]defStructs.Champ{
	348: {"AdvisedChargeCurrency", -1, 2, nil},
	356: {"TotalAdvisedCharge", -1, 1, &TotalValueT13},
	357: {"TotalAdvisedChargeRefund", -1, 1, &TotalValueT13},
	358: {"TotalCommission", -1, 1, &TotalValueT13},
	359: {"TotalCommissionRefund", -1, 1, &TotalValueT13},
}

var TotChargValT12 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	241: {"AbsoluteAmount", -1, 1, nil},
	355: {"TotalChargeRefund", -1, 1, &TotalValueT13},
}

var ValAddServUsedT12 = map[int]defStructs.Champ{
	69:  {"ChargeInformation", -1, 1, &ChargInfoT14},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT13},
	237: {"VasCode", -1, 1, &RecEntCodeT13},
}

var SupplServUsedT12 = map[int]defStructs.Champ{
	37:  {"BasicServiceCodeList", -1, 1, &BasicServCodeListT13},
	69:  {"ChargeInformation", -1, 1, &ChargInfoT14},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT13},
	204: {"SsParameters", -1, 2, nil},
	208: {"SupplServiceActionCod", -1, 1, nil},
	209: {"SupplServiceCode", -1, 2, nil},
	219: {"ThirdPartyInformation", -1, 1, &CallOriginT13},
}

var CallOriginT12 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	75: {"ClirIndicator", -1, 1, nil},
}

var MinChargSubT12 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var SimChargSubT12 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var RapFileSeqNumberT12 = map[int]defStructs.Champ{
	109: {"FileSequenceNumber", -1, 2, nil},
}

var DestNetworkT12 = map[int]defStructs.Champ{
	155: {"NetworkId", -1, 2, nil},
}

var DestinT12 = map[int]defStructs.Champ{
	33:  {"AddressStringDigits", -1, 1, nil},
	42:  {"CalledPlace", -1, 2, nil},
	46:  {"CalledRegion", -1, 2, nil},
	247: {"CalledNumAnalysisCode", -1, 1, &RecEntCodeT13},
	279: {"DialledDigits", -1, 2, nil},
}

var NetworkLocT12 = map[int]defStructs.Champ{
	45:  {"CallReference", -1, 1, nil},
	59:  {"CellId", -1, 0, nil},
	136: {"LocationArea", -1, 0, nil},
	184: {"RecEntityCode", -1, 1, &RecEntCodeT13},
}

var HomeLocInfoT12 = map[int]defStructs.Champ{
	122: {"HomeBid", -1, 1, &ServBidT13},
	137: {"LocationDescription", -1, 2, nil},
}

var GeoLocT12 = map[int]defStructs.Champ{
	134: {"Latitude", -1, 1, nil},
	137: {"LocationDescription", -1, 2, nil},
	139: {"Longitude", -1, 1, nil},
	195: {"ServingNetwork", -1, 2, nil},
	198: {"ServingBid", -1, 1, &ServBidT13},
}

var CamServAddT12 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
}

var GprsDestT12 = map[int]defStructs.Champ{
	190: {"RemotePdpAddressList", -1, 1, &RemPdpAddListT13},
	261: {"AccessPointNameNI", -1, 2, nil},
	262: {"AccessPointNameOI", -1, 2, nil},
}

var CamModifListT12 = map[int]defStructs.Champ{
	53: {"CamelModification", -1, 1, nil},
}

var CamCallRefT12 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	50: {"CamelCallReferenceNumber", -1, 2, nil},
}

var BasServUsedT12 = map[int]defStructs.Champ{
	36:  {"BasicService", -1, 1, &BasiCServT13},
	70:  {"ChargeInformationList", -1, 1, &ChargInfoListT13},
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT13},
	82:  {"DataVolumeReference", -1, 1, nil},
	124: {"HSCSDInformation", -1, 1, &HSCSDinfoT13},
}

var TotalValueT12 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var RecEntCodeT12 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var TaxInfoT12 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, nil},
	398: {"TaxableAmount", -1, 1, nil},
}

var DateT12 = map[int]defStructs.Champ{
	83: {"DateTime", -1, 1, &DateTimeT13},
}

var GprsServUsT13 = map[int]defStructs.Champ{
	74:  {"ChargingTimeStamp", -1, 1, &DateTimeT14},
	250: {"DataVolumeIncoming", -1, 0, &DataVolT14},
	251: {"DataVolumeOutgoing", -1, 0, &DataVolT14},
	270: {"UMTSQoSRequested", -1, 1, &UMTSQoST14},
	271: {"UMTSQoSUsed", -1, 1, &UMTSQoST14},
	272: {"GSMQoSRequested", -1, 1, &GSMQosT14},
	273: {"GSMQoSUsed", -1, 1, &GSMQosT14},
}

var RecEntCodeListT13 = map[int]defStructs.Champ{
	184: {"RecEntityCode", -1, 1, &RecEntCodeT14},
}

var MinChargSubT13 = map[int]defStructs.Champ{
	146: {"Min", -1, 1, nil},
	253: {"Mdn", -1, 2, nil},
}

var SimChargSubT13 = map[int]defStructs.Champ{
	129: {"Imsi", -1, 0, nil},
	152: {"Msisdn", -1, 0, nil},
}

var PdpAddT13 = map[int]defStructs.Champ{
	165: {"PacketDataProtocolAddress", -1, 2, nil},
}

var TotalValueT13 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var CallOriginT13 = map[int]defStructs.Champ{
	33: {"AddressStringDigits", -1, 1, nil},
	75: {"ClirIndicator", -1, 1, nil},
}

var BasicServCodeListT13 = map[int]defStructs.Champ{
	40:  {"BearerServiceCode", -1, 2, nil},
	218: {"TeleServiceCode", -1, 2, nil},
}

var RecEntCodeT13 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var ServBidT13 = map[int]defStructs.Champ{
	242: {"Bid", -1, 2, nil},
}

var RemPdpAddListT13 = map[int]defStructs.Champ{
	165: {"PacketDataProtocolAddress", -1, 2, nil},
}

var DateTimeT13 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT14},
}

var BasiCServT13 = map[int]defStructs.Champ{
	34:  {"AiurRequested", -1, 1, nil},
	40:  {"BearerServiceCode", -1, 2, nil},
	111: {"Fnur", -1, 1, nil},
	179: {"RadioChannetRequested", -1, 1, nil},
	180: {"RadioChannelUsed", -1, 1, nil},
	202: {"SpeedVersionRequested", -1, 1, nil},
	203: {"SpeedVersionUsed", -1, 1, nil},
	218: {"TeleServiceCode", -1, 2, nil},
	228: {"TransparencyIndicator", -1, 1, nil},
	280: {"UserProtocolIndicator", -1, 1, nil},
}

var ChargInfoListT13 = map[int]defStructs.Champ{
	69: {"ChargeInformation", -1, 1, &ChargInfoT14},
}

var HSCSDinfoT13 = map[int]defStructs.Champ{
	35:  {"BasicHSCSDParameters", -1, 1, &BasHSCDParamT14},
	140: {"HSCSDParameterModificationList", -1, 1, &HSCSDParamModListT14},
}

var GSMQosT14 = map[int]defStructs.Champ{
	171: {"QoSDelay", -1, 1, nil},
	173: {"QoSMeanThrougtput", -1, 1, nil},
	174: {"QoSPeakThrougtput", -1, 1, nil},
	175: {"QoSPrecedence", -1, 1, nil},
	176: {"QoSreliability", -1, 1, nil},
}

var UMTSQoST14 = map[int]defStructs.Champ{
	264: {"QoSMaxBitRateUplink", -1, 2, nil},
	268: {"QoSTrafficClass", -1, 1, nil},
	274: {"QoSMaxBitRateDownlink", -1, 1, nil},
	275: {"QoSAlloRetenPriority", -1, 1, nil},
	283: {"QoSGuaranteedBitRateDownlink", -1, 2, nil},
	284: {"QoSGuaranteedBitRateUplink", -1, 2, nil},
}

var DataVolT14 = map[int]defStructs.Champ{
	81: {"DataVolume", -1, 1, nil},
}

var DateTimeT14 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT15},
}

var BasHSCDParamT14 = map[int]defStructs.Champ{
	60:  {"ChannelCoding", -1, 1, nil},
	61:  {"ChallelAcceptableList", -1, 1, &ChannelCodT15},
	161: {"NumberOfChannels", -1, 1, nil},
	252: {"NumberOfChannelUsed", -1, 1, nil},
}

var RecEntCodeT14 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var ChargInfoT14 = map[int]defStructs.Champ{
	66:  {"ChargeItem", -1, 2, nil},
	64:  {"ChargeDetList", -1, 1, &ChargDetListT15},
	96:  {"DiscountInformation", -1, 1, &DiscInfoT15},
	105: {"Exchange", -1, 1, &RecEntCodeT15},
	214: {"TaxInformationList", -1, 1, &TaxInfoListT15},
	258: {"CallTypeGroup", -1, 1, &CallTypeGrT15},
}

var HSCSDParamModListT14 = map[int]defStructs.Champ{
	125: {"HSCSDParameterModification", -1, 1, &HSCSDParamModT15},
}

var ChargDetListT15 = map[int]defStructs.Champ{
	63: {"ChargeDetail", -1, 1, &ChargDetT16},
}

var RecEntCodeT15 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}

var HSCSDParamModT15 = map[int]defStructs.Champ{
	34:  {"AiurRequested", -1, 1, nil},
	60:  {"ChannelCoding", -1, 1, nil},
	130: {"InitiatingParty", -1, 1, nil},
	150: {"ModificationTimeStamp", -1, 1, &DateT16},
	161: {"NumberOfChannels", -1, 1, nil},
	252: {"NumberOfChannelUsed", -1, 1, nil},
}

var ChannelCodT15 = map[int]defStructs.Champ{
	60: {"ChannelCoding", -1, 1, nil},
}

var ChargDetT15 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT16},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, nil},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var DiscInfoT15 = map[int]defStructs.Champ{
	91: {"DiscountCode", -1, 1, nil},
	93: {"DiscountValue", -1, 1, nil},
}

var TaxInfoListT15 = map[int]defStructs.Champ{
	9:   {"MobileOriginatedCall", -1, 1, nil},
	10:  {"MobileTerminatedCall", -1, 1, nil},
	14:  {"GprsCall", -1, 1, nil},
	15:  {"AuditControlInfo", -1, 1, nil},
	33:  {"AddressStringDigits", -1, 1, nil},
	162: {"OperatorSpecInfoList", -1, 1, nil},
	213: {"TaxInformation", -1, 1, &TaxInfoT16},
}

var CallTypeGrT15 = map[int]defStructs.Champ{
	255: {"CallTypeLevel2", -1, 1, nil},
	256: {"CallTypeLevel3", -1, 1, nil},
	257: {"CalledCountryCode", -1, 1, nil},
	259: {"CallTypeLevel1", -1, 1, nil},
}

var ChargDetT16 = map[int]defStructs.Champ{
	62:  {"Charge", -1, 1, &TotalValueT17},
	65:  {"ChargeableUnits", -1, 0, nil},
	68:  {"ChargedUnits", -1, 0, nil},
	71:  {"ChargeType", -1, 2, nil},
	74:  {"ChargingTimeStamp", -1, 1, nil},
	85:  {"DayCategory", -1, 2, nil},
	86:  {"DayCategorySubType", -1, 2, nil},
	154: {"MultiRateIndicator", -1, 1, nil},
	220: {"TimeBand", -1, 2, nil},
	221: {"TimeBandSubType", -1, 2, nil},
}

var TaxInfoT16 = map[int]defStructs.Champ{
	71:  {"ChargeType", -1, 2, nil},
	212: {"TaxCode", -1, 1, nil},
	397: {"TaxValue", -1, 1, nil},
	398: {"TaxableAmount", -1, 1, nil},
}

var TotalValueT16 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateT16 = map[int]defStructs.Champ{
	83: {"DateTime", -1, 1, &DateTimeT17},
}

var TotalValueT17 = map[int]defStructs.Champ{
	241: {"AbsoluteAmount", -1, 1, nil},
}

var DateTimeT17 = map[int]defStructs.Champ{
	16:  {"LocalTimeStamp", -1, 2, nil},
	232: {"UtcTimeOffsetCode", -1, 1, &RecEntCodeT18},
}

var RecEntCodeT18 = map[int]defStructs.Champ{
	243: {"Code", -1, 1, nil},
}
