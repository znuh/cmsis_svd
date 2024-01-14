// Code generated by xgen. DO NOT EDIT.

package cmsis_svd

import (
	"encoding/xml"
	"os"
)

type Range struct {
	XMLName xml.Name `xml:"range"`
	Minimum string   `xml:"minimum"`
	Maximum string   `xml:"maximum"`
}

type WriteConstraint struct {
	XMLName             xml.Name `xml:"writeConstraint"`
	WriteAsRead         bool     `xml:"writeAsRead"`
	UseEnumeratedValues bool     `xml:"useEnumeratedValues"`
	Range               *Range   `xml:"range"`
}

type AddressBlock struct {
	XMLName    xml.Name `xml:"addressBlock"`
	Offset     string   `xml:"offset"`
	Size       string   `xml:"size"`
	Usage      string   `xml:"usage"`
	Protection string   `xml:"protection"`
}

type Interrupt struct {
	XMLName     xml.Name `xml:"interrupt"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Value       int      `xml:"value"`
}

type RegisterPropertiesGroup struct {
	XMLName    xml.Name `xml:"registerPropertiesGroup"`
	Size       string
	Access     string
	Protection string
	ResetValue string
	ResetMask  string
}

type BitRangeLsbMsbStyle struct {
	XMLName xml.Name `xml:"bitRangeLsbMsbStyle"`
	Lsb     string
	Msb     string
}

type BitRangeOffsetWidthStyle struct {
	XMLName   xml.Name `xml:"bitRangeOffsetWidthStyle"`
	BitOffset string
	BitWidth  string
}

type DimElementGroup struct {
	XMLName      xml.Name `xml:"dimElementGroup"`
	Dim          string
	DimIncrement string
	DimIndex     string
}

type Region struct {
	XMLName     xml.Name `xml:"region"`
	EnabledAttr bool     `xml:"enabled,attr,omitempty"`
	NameAttr    string   `xml:"name,attr,omitempty"`
	Base        string   `xml:"base"`
	Limit       string   `xml:"limit"`
	Access      string   `xml:"access"`
}

type SauRegionsConfig struct {
	XMLName                    xml.Name  `xml:"sauRegionsConfig"`
	EnabledAttr                bool      `xml:"enabled,attr,omitempty"`
	ProtectionWhenDisabledAttr string    `xml:"protectionWhenDisabled,attr,omitempty"`
	Region                     []*Region `xml:"region"`
}

type Cpu struct {
	XMLName             xml.Name          `xml:"cpu"`
	Name                string            `xml:"name"`
	Revision            string            `xml:"revision"`
	Endian              string            `xml:"endian"`
	MpuPresent          bool              `xml:"mpuPresent"`
	FpuPresent          bool              `xml:"fpuPresent"`
	FpuDP               bool              `xml:"fpuDP"`
	IcachePresent       bool              `xml:"icachePresent"`
	DcachePresent       bool              `xml:"dcachePresent"`
	ItcmPresent         bool              `xml:"itcmPresent"`
	DtcmPresent         bool              `xml:"dtcmPresent"`
	VtorPresent         bool              `xml:"vtorPresent"`
	NvicPrioBits        string            `xml:"nvicPrioBits"`
	VendorSystickConfig bool              `xml:"vendorSystickConfig"`
	DeviceNumInterrupts string            `xml:"deviceNumInterrupts"`
	SauNumRegions       string            `xml:"sauNumRegions"`
	SauRegionsConfig    *SauRegionsConfig `xml:"sauRegionsConfig"`
}

type EnumeratedValue struct {
	XMLName     xml.Name `xml:"enumeratedValue"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Value       string   `xml:"value"`
	IsDefault   bool     `xml:"isDefault"`
}

type EnumeratedValues struct {
	XMLName         xml.Name           `xml:"enumeratedValues"`
	DerivedFromAttr string             `xml:"derivedFrom,attr,omitempty"`
	Name            string             `xml:"name"`
	Usage           string             `xml:"usage"`
	EnumeratedValue []*EnumeratedValue `xml:"enumeratedValue"`
}

type Field struct {
	XMLName                  xml.Name `xml:"field"`
	DerivedFromAttr          string   `xml:"derivedFrom,attr,omitempty"`
	BitRangeLsbMsbStyle      *BitRangeLsbMsbStyle
	BitRangeOffsetWidthStyle *BitRangeOffsetWidthStyle
	Name                     string                  `xml:"name"`
	Description              string                  `xml:"description"`
	BitRange                 string                  `xml:"bitRange"`
	Access                   string                  `xml:"access"`
	ModifiedWriteValues      string                  `xml:"modifiedWriteValues"`
	WriteConstraint          *WriteConstraint    `xml:"writeConstraint"`
	ReadAction               string                  `xml:"readAction"`
	EnumeratedValues         []*EnumeratedValues `xml:"enumeratedValues"`
}

type Fields struct {
	XMLName xml.Name     `xml:"fields"`
	Field   []*Field `xml:"field"`
}

type Register struct {
	XMLName                 xml.Name `xml:"register"`
	DerivedFromAttr         string   `xml:"derivedFrom,attr,omitempty"`
	DimElementGroup         *DimElementGroup
	RegisterPropertiesGroup *RegisterPropertiesGroup
	Name                    string               `xml:"name"`
	DisplayName             string               `xml:"displayName"`
	Description             string               `xml:"description"`
	AlternateGroup          string               `xml:"alternateGroup"`
	AlternateRegister       string               `xml:"alternateRegister"`
	AddressOffset           string               `xml:"addressOffset"`
	DataType                string               `xml:"dataType"`
	ModifiedWriteValues     string               `xml:"modifiedWriteValues"`
	WriteConstraint         *WriteConstraint `xml:"writeConstraint"`
	ReadAction              string               `xml:"readAction"`
	Fields                  *Fields          `xml:"fields"`
}

type Cluster struct {
	XMLName                 xml.Name `xml:"cluster"`
	DerivedFromAttr         string   `xml:"derivedFrom,attr,omitempty"`
	DimElementGroup         *DimElementGroup
	RegisterPropertiesGroup *RegisterPropertiesGroup
	Name                    string          `xml:"name"`
	Description             string          `xml:"description"`
	AlternateCluster        string          `xml:"alternateCluster"`
	HeaderStructName        string          `xml:"headerStructName"`
	AddressOffset           string          `xml:"addressOffset"`
	Register                []*Register `xml:"register"`
	Cluster                 []*Cluster  `xml:"cluster"`
}

type Registers struct {
	XMLName  xml.Name        `xml:"registers"`
	Cluster  []*Cluster  `xml:"cluster"`
	Register []*Register `xml:"register"`
}

type Peripheral struct {
	XMLName                 xml.Name `xml:"peripheral"`
	DerivedFromAttr         string   `xml:"derivedFrom,attr,omitempty"`
	DimElementGroup         *DimElementGroup
	RegisterPropertiesGroup *RegisterPropertiesGroup
	Name                    string              `xml:"name"`
	Version                 string              `xml:"version"`
	Description             string              `xml:"description"`
	AlternatePeripheral     string              `xml:"alternatePeripheral"`
	GroupName               string              `xml:"groupName"`
	PrependToName           string              `xml:"prependToName"`
	AppendToName            string              `xml:"appendToName"`
	HeaderStructName        string              `xml:"headerStructName"`
	DisableCondition        string              `xml:"disableCondition"`
	BaseAddress             string              `xml:"baseAddress"`
	AddressBlock            []*AddressBlock `xml:"addressBlock"`
	Interrupt               []*Interrupt    `xml:"interrupt"`
	Registers               *Registers      `xml:"registers"`
}

type Peripherals struct {
	XMLName    xml.Name          `xml:"peripherals"`
	Peripheral []*Peripheral `xml:"peripheral"`
}

type VendorExtensions struct {
	XMLName xml.Name `xml:"vendorExtensions"`
}

type Device struct {
	XMLName                 xml.Name `xml:"device"`
	SchemaVersionAttr       float64  `xml:"schemaVersion,attr"`
	RegisterPropertiesGroup *RegisterPropertiesGroup
	Vendor                  string            `xml:"vendor"`
	VendorID                string            `xml:"vendorID"`
	Name                    string            `xml:"name"`
	Series                  string            `xml:"series"`
	Version                 string            `xml:"version"`
	Description             string            `xml:"description"`
	LicenseText             string            `xml:"licenseText"`
	Cpu                     *Cpu              `xml:"cpu"`
	HeaderSystemFilename    string            `xml:"headerSystemFilename"`
	HeaderDefinitionsPrefix string            `xml:"headerDefinitionsPrefix"`
	AddressUnitBits         string            `xml:"addressUnitBits"`
	Width                   string            `xml:"width"`
	Peripherals             *Peripherals      `xml:"peripherals"`
	VendorExtensions        *VendorExtensions `xml:"vendorExtensions"`
}

func ReadSVD(fname string) (*Device, error) {
	f, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    
    device  := &Device{}
    decoder := xml.NewDecoder(f)
    if err := decoder.Decode(device); err != nil {
		return nil, err
	}
	return device, nil
}
