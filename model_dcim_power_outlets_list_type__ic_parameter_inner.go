/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.3 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// DcimPowerOutletsListTypeIcParameterInner the model 'DcimPowerOutletsListTypeIcParameterInner'
type DcimPowerOutletsListTypeIcParameterInner string

// List of dcim_power_outlets_list_type__ic_parameter_inner
const (
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_EMPTY DcimPowerOutletsListTypeIcParameterInner = ""
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_CS6360_C DcimPowerOutletsListTypeIcParameterInner = "CS6360C"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_CS6364_C DcimPowerOutletsListTypeIcParameterInner = "CS6364C"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_CS8164_C DcimPowerOutletsListTypeIcParameterInner = "CS8164C"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_CS8264_C DcimPowerOutletsListTypeIcParameterInner = "CS8264C"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_CS8364_C DcimPowerOutletsListTypeIcParameterInner = "CS8364C"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_CS8464_C DcimPowerOutletsListTypeIcParameterInner = "CS8464C"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_DC_TERMINAL DcimPowerOutletsListTypeIcParameterInner = "dc-terminal"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_EATON_C39 DcimPowerOutletsListTypeIcParameterInner = "eaton-c39"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_HARDWIRED DcimPowerOutletsListTypeIcParameterInner = "hardwired"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_HDOT_CX DcimPowerOutletsListTypeIcParameterInner = "hdot-cx"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_2P_E_4H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-2p-e-4h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_2P_E_6H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-2p-e-6h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_2P_E_9H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-2p-e-9h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_3P_E_4H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-3p-e-4h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_3P_E_6H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-3p-e-6h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_3P_E_9H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-3p-e-9h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_3P_N_E_4H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-3p-n-e-4h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_3P_N_E_6H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-3p-n-e-6h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_3P_N_E_9H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-3p-n-e-9h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_P_N_E_4H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-p-n-e-4h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_P_N_E_6H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-p-n-e-6h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60309_P_N_E_9H DcimPowerOutletsListTypeIcParameterInner = "iec-60309-p-n-e-9h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60320_C13 DcimPowerOutletsListTypeIcParameterInner = "iec-60320-c13"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60320_C15 DcimPowerOutletsListTypeIcParameterInner = "iec-60320-c15"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60320_C19 DcimPowerOutletsListTypeIcParameterInner = "iec-60320-c19"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60320_C21 DcimPowerOutletsListTypeIcParameterInner = "iec-60320-c21"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60320_C5 DcimPowerOutletsListTypeIcParameterInner = "iec-60320-c5"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60320_C7 DcimPowerOutletsListTypeIcParameterInner = "iec-60320-c7"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_IEC_60906_1 DcimPowerOutletsListTypeIcParameterInner = "iec-60906-1"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_E DcimPowerOutletsListTypeIcParameterInner = "ita-e"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_F DcimPowerOutletsListTypeIcParameterInner = "ita-f"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_G DcimPowerOutletsListTypeIcParameterInner = "ita-g"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_H DcimPowerOutletsListTypeIcParameterInner = "ita-h"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_I DcimPowerOutletsListTypeIcParameterInner = "ita-i"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_J DcimPowerOutletsListTypeIcParameterInner = "ita-j"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_K DcimPowerOutletsListTypeIcParameterInner = "ita-k"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_L DcimPowerOutletsListTypeIcParameterInner = "ita-l"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_M DcimPowerOutletsListTypeIcParameterInner = "ita-m"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_MULTISTANDARD DcimPowerOutletsListTypeIcParameterInner = "ita-multistandard"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_N DcimPowerOutletsListTypeIcParameterInner = "ita-n"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_ITA_O DcimPowerOutletsListTypeIcParameterInner = "ita-o"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_MOLEX_MICRO_FIT_1X2 DcimPowerOutletsListTypeIcParameterInner = "molex-micro-fit-1x2"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_MOLEX_MICRO_FIT_2X2 DcimPowerOutletsListTypeIcParameterInner = "molex-micro-fit-2x2"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_MOLEX_MICRO_FIT_2X4 DcimPowerOutletsListTypeIcParameterInner = "molex-micro-fit-2x4"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NBR_14136_10A DcimPowerOutletsListTypeIcParameterInner = "nbr-14136-10a"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NBR_14136_20A DcimPowerOutletsListTypeIcParameterInner = "nbr-14136-20a"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_1_15R DcimPowerOutletsListTypeIcParameterInner = "nema-1-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_10_30R DcimPowerOutletsListTypeIcParameterInner = "nema-10-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_10_50R DcimPowerOutletsListTypeIcParameterInner = "nema-10-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_14_20R DcimPowerOutletsListTypeIcParameterInner = "nema-14-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_14_30R DcimPowerOutletsListTypeIcParameterInner = "nema-14-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_14_50R DcimPowerOutletsListTypeIcParameterInner = "nema-14-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_14_60R DcimPowerOutletsListTypeIcParameterInner = "nema-14-60r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_15_15R DcimPowerOutletsListTypeIcParameterInner = "nema-15-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_15_20R DcimPowerOutletsListTypeIcParameterInner = "nema-15-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_15_30R DcimPowerOutletsListTypeIcParameterInner = "nema-15-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_15_50R DcimPowerOutletsListTypeIcParameterInner = "nema-15-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_15_60R DcimPowerOutletsListTypeIcParameterInner = "nema-15-60r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_5_15R DcimPowerOutletsListTypeIcParameterInner = "nema-5-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_5_20R DcimPowerOutletsListTypeIcParameterInner = "nema-5-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_5_30R DcimPowerOutletsListTypeIcParameterInner = "nema-5-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_5_50R DcimPowerOutletsListTypeIcParameterInner = "nema-5-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_6_15R DcimPowerOutletsListTypeIcParameterInner = "nema-6-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_6_20R DcimPowerOutletsListTypeIcParameterInner = "nema-6-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_6_30R DcimPowerOutletsListTypeIcParameterInner = "nema-6-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_6_50R DcimPowerOutletsListTypeIcParameterInner = "nema-6-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L1_15R DcimPowerOutletsListTypeIcParameterInner = "nema-l1-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L10_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l10-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L14_20R DcimPowerOutletsListTypeIcParameterInner = "nema-l14-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L14_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l14-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L14_50R DcimPowerOutletsListTypeIcParameterInner = "nema-l14-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L14_60R DcimPowerOutletsListTypeIcParameterInner = "nema-l14-60r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L15_20R DcimPowerOutletsListTypeIcParameterInner = "nema-l15-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L15_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l15-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L15_50R DcimPowerOutletsListTypeIcParameterInner = "nema-l15-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L15_60R DcimPowerOutletsListTypeIcParameterInner = "nema-l15-60r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L21_20R DcimPowerOutletsListTypeIcParameterInner = "nema-l21-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L21_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l21-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L22_20R DcimPowerOutletsListTypeIcParameterInner = "nema-l22-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L22_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l22-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L5_15R DcimPowerOutletsListTypeIcParameterInner = "nema-l5-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L5_20R DcimPowerOutletsListTypeIcParameterInner = "nema-l5-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L5_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l5-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L5_50R DcimPowerOutletsListTypeIcParameterInner = "nema-l5-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L6_15R DcimPowerOutletsListTypeIcParameterInner = "nema-l6-15r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L6_20R DcimPowerOutletsListTypeIcParameterInner = "nema-l6-20r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L6_30R DcimPowerOutletsListTypeIcParameterInner = "nema-l6-30r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEMA_L6_50R DcimPowerOutletsListTypeIcParameterInner = "nema-l6-50r"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEUTRIK_POWERCON_20A DcimPowerOutletsListTypeIcParameterInner = "neutrik-powercon-20a"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEUTRIK_POWERCON_32A DcimPowerOutletsListTypeIcParameterInner = "neutrik-powercon-32a"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEUTRIK_POWERCON_TRUE1 DcimPowerOutletsListTypeIcParameterInner = "neutrik-powercon-true1"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_NEUTRIK_POWERCON_TRUE1_TOP DcimPowerOutletsListTypeIcParameterInner = "neutrik-powercon-true1-top"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_OTHER DcimPowerOutletsListTypeIcParameterInner = "other"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_SAF_D_GRID DcimPowerOutletsListTypeIcParameterInner = "saf-d-grid"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_UBIQUITI_SMARTPOWER DcimPowerOutletsListTypeIcParameterInner = "ubiquiti-smartpower"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_USB_A DcimPowerOutletsListTypeIcParameterInner = "usb-a"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_USB_C DcimPowerOutletsListTypeIcParameterInner = "usb-c"
	DCIMPOWEROUTLETSLISTTYPEICPARAMETERINNER_USB_MICRO_B DcimPowerOutletsListTypeIcParameterInner = "usb-micro-b"
)

// All allowed values of DcimPowerOutletsListTypeIcParameterInner enum
var AllowedDcimPowerOutletsListTypeIcParameterInnerEnumValues = []DcimPowerOutletsListTypeIcParameterInner{
	"",
	"CS6360C",
	"CS6364C",
	"CS8164C",
	"CS8264C",
	"CS8364C",
	"CS8464C",
	"dc-terminal",
	"eaton-c39",
	"hardwired",
	"hdot-cx",
	"iec-60309-2p-e-4h",
	"iec-60309-2p-e-6h",
	"iec-60309-2p-e-9h",
	"iec-60309-3p-e-4h",
	"iec-60309-3p-e-6h",
	"iec-60309-3p-e-9h",
	"iec-60309-3p-n-e-4h",
	"iec-60309-3p-n-e-6h",
	"iec-60309-3p-n-e-9h",
	"iec-60309-p-n-e-4h",
	"iec-60309-p-n-e-6h",
	"iec-60309-p-n-e-9h",
	"iec-60320-c13",
	"iec-60320-c15",
	"iec-60320-c19",
	"iec-60320-c21",
	"iec-60320-c5",
	"iec-60320-c7",
	"iec-60906-1",
	"ita-e",
	"ita-f",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-multistandard",
	"ita-n",
	"ita-o",
	"molex-micro-fit-1x2",
	"molex-micro-fit-2x2",
	"molex-micro-fit-2x4",
	"nbr-14136-10a",
	"nbr-14136-20a",
	"nema-1-15r",
	"nema-10-30r",
	"nema-10-50r",
	"nema-14-20r",
	"nema-14-30r",
	"nema-14-50r",
	"nema-14-60r",
	"nema-15-15r",
	"nema-15-20r",
	"nema-15-30r",
	"nema-15-50r",
	"nema-15-60r",
	"nema-5-15r",
	"nema-5-20r",
	"nema-5-30r",
	"nema-5-50r",
	"nema-6-15r",
	"nema-6-20r",
	"nema-6-30r",
	"nema-6-50r",
	"nema-l1-15r",
	"nema-l10-30r",
	"nema-l14-20r",
	"nema-l14-30r",
	"nema-l14-50r",
	"nema-l14-60r",
	"nema-l15-20r",
	"nema-l15-30r",
	"nema-l15-50r",
	"nema-l15-60r",
	"nema-l21-20r",
	"nema-l21-30r",
	"nema-l22-20r",
	"nema-l22-30r",
	"nema-l5-15r",
	"nema-l5-20r",
	"nema-l5-30r",
	"nema-l5-50r",
	"nema-l6-15r",
	"nema-l6-20r",
	"nema-l6-30r",
	"nema-l6-50r",
	"neutrik-powercon-20a",
	"neutrik-powercon-32a",
	"neutrik-powercon-true1",
	"neutrik-powercon-true1-top",
	"other",
	"saf-d-grid",
	"ubiquiti-smartpower",
	"usb-a",
	"usb-c",
	"usb-micro-b",
}

func (v *DcimPowerOutletsListTypeIcParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimPowerOutletsListTypeIcParameterInner(value)
	for _, existing := range AllowedDcimPowerOutletsListTypeIcParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimPowerOutletsListTypeIcParameterInner", value)
}

// NewDcimPowerOutletsListTypeIcParameterInnerFromValue returns a pointer to a valid DcimPowerOutletsListTypeIcParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimPowerOutletsListTypeIcParameterInnerFromValue(v string) (*DcimPowerOutletsListTypeIcParameterInner, error) {
	ev := DcimPowerOutletsListTypeIcParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimPowerOutletsListTypeIcParameterInner: valid values are %v", v, AllowedDcimPowerOutletsListTypeIcParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimPowerOutletsListTypeIcParameterInner) IsValid() bool {
	for _, existing := range AllowedDcimPowerOutletsListTypeIcParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_power_outlets_list_type__ic_parameter_inner value
func (v DcimPowerOutletsListTypeIcParameterInner) Ptr() *DcimPowerOutletsListTypeIcParameterInner {
	return &v
}

type NullableDcimPowerOutletsListTypeIcParameterInner struct {
	value *DcimPowerOutletsListTypeIcParameterInner
	isSet bool
}

func (v NullableDcimPowerOutletsListTypeIcParameterInner) Get() *DcimPowerOutletsListTypeIcParameterInner {
	return v.value
}

func (v *NullableDcimPowerOutletsListTypeIcParameterInner) Set(val *DcimPowerOutletsListTypeIcParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimPowerOutletsListTypeIcParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimPowerOutletsListTypeIcParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimPowerOutletsListTypeIcParameterInner(val *DcimPowerOutletsListTypeIcParameterInner) *NullableDcimPowerOutletsListTypeIcParameterInner {
	return &NullableDcimPowerOutletsListTypeIcParameterInner{value: val, isSet: true}
}

func (v NullableDcimPowerOutletsListTypeIcParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimPowerOutletsListTypeIcParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

