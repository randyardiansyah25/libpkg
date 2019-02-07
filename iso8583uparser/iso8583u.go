package iso8583uParser

import (
	"errors"
	"fmt"
	"github.com/mofax/iso8583"
	"github.com/randyardiansyah25/libpkg/util/str"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const DefaultSpecFile string = "isopackager.yml"

/*
Params : berupa informasi filename (Optional)
jika pemanggilan function tidak mengirimkan param, filename akan mengacu pada default_specfile
*/
func NewISO8583U(params ...string) (ISO8583U, error) {
	var iso ISO8583U
	if len(params) > 0 {
		iso.SpecFile = params[0]
	} else {
		iso.SpecFile = DefaultSpecFile
	}
	err := iso.prepare()
	return iso, err
}

type FieldDescription struct {
	ContentType string `yaml:"ContentType"`
	MaxLen      int    `yaml:"MaxLen"`
	MinLen      int    `yaml:"MinLen"`
	LenType     string `yaml:"LenType"`
	Label       string `yaml:"Label"`
}

type ISO8583U struct {
	SpecFile   string
	isoMTI     string
	isoBitmap  string
	isoElement map[int64]string
	fields     map[int]FieldDescription
}

func (p *ISO8583U) prepare() error {
	p.isoElement = make(map[int64]string, 0)
	ymlContent, err := ioutil.ReadFile(p.getSpecFile())
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(ymlContent, &p.fields)
	return err
}

func (p *ISO8583U) getSpecFile() string {
	if len(p.SpecFile) == 0 {
		return DefaultSpecFile
	}
	return p.SpecFile
}

func (p *ISO8583U) GoUnMarshal(message string) error {
	isoFormatter := iso8583.NewISOStruct(p.getSpecFile(), true)
	isoStructParser, err := isoFormatter.Parse(message)
	if err != nil {
		return errors.New(fmt.Sprint("Parse iso message failed, cause : ", err.Error()))
	}
	p.isoMTI = isoStructParser.Mti.String()
	p.isoBitmap, _ = iso8583.BitMapArrayToHex(isoStructParser.Bitmap)
	p.isoElement = isoStructParser.Elements.GetElements()
	return nil
}

func (p *ISO8583U) GetField(fieldNo int64) string {
	return strings.TrimSpace(p.isoElement[fieldNo])
}

func (p *ISO8583U) GetBitmap() string {
	return p.isoBitmap
}

func (p *ISO8583U) SetMti(mti string) {
	p.isoMTI = mti
}

func (p *ISO8583U) SetField(fieldNo int64, value interface{}) {
	sValue := fmt.Sprint(value)

	fieldDef := p.fields[int(fieldNo)]
	if fieldDef.LenType == "fixed" {
		if fieldDef.ContentType == "n" {
			sValue = strutils.LeftPad(sValue, fieldDef.MaxLen, "0")
		} else {
			sValue = strutils.RightPad(sValue, fieldDef.MaxLen, " ")
		}
	}
	p.isoElement[fieldNo] = sValue
}

func (p *ISO8583U) GoMarshal() (string, error) {
	isoFormatter := iso8583.NewISOStruct(p.getSpecFile(), true)
	if len(p.isoMTI) <= 0 {
		return "", errors.New("Empty generates invalid MTI")
	}

	_ = isoFormatter.AddMTI(p.isoMTI)

	for field, val := range p.isoElement {
		_ = isoFormatter.AddField(field, val)
	}
	isoMsgStr, err := isoFormatter.ToString()
	if err != nil {
		return "", errors.New("Cannot marshal iso format, cause : " + err.Error())
	}
	p.isoBitmap, _ = iso8583.BitMapArrayToHex(isoFormatter.Bitmap)
	return isoMsgStr, nil
}

func (p *ISO8583U) PrettyPrint() string {

	isoBuffer := []string{
		fmt.Sprintf("[%s][%s]\n", strutils.LeftPad("1", 3, "0"), p.isoMTI),
		fmt.Sprintf("[%s][%s]\n", strutils.LeftPad("2", 3, "0"), p.isoBitmap),
	}

	keys := make([]int, 0)
	for k := range p.isoElement {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		isoBuffer = append(isoBuffer, fmt.Sprintf("[%s][%s]\n", strutils.LeftPad(strconv.Itoa(k), 3, "0"), p.isoElement[int64(k)]))
	}
	return strings.Join(isoBuffer, "")
}
