package payment

import (
	"github.com/abyssparanoia/go-gmo/internal/pkg/parser"
	"github.com/abyssparanoia/go-gmo/internal/pkg/validate"
)

// SaveCardRequest ... save card request
type SaveCardRequest struct {
	MemberID     string `schema:"MemberID" validate:"required,max=60"`
	SeqMode      string `schema:"SeqMode"`
	CardSeq      string `schema:"CardSeq"`
	DefaultFlag  string `schema:"DefaultFlag"`
	CardName     string `schema:"CardName"`
	CardNo       string `schema:"CardNo" validate:"required,len=16"`
	CardPass     string `schema:"CardPass"`
	Expire       string `schema:"Expire" validate:"required,len=4"`
	HolderName   string `schema:"HolderName"`
	Token        string `schema:"Token"`
	UpdateType   string `schema:"UpdateType"`
	SecurityCode string `schema:"SecurityCode"`
}

// Validate ... validate
func (r *SaveCardRequest) Validate() error {
	return validate.Struct(r)
}

// SaveCardResponse ... save card response
type SaveCardResponse struct {
	CardSeq                string `schema:"CardSeq"`
	CardNo                 string `schema:"CardNo"`
	Forward                string `schema:"Forward"`
	ErrCode                string `schema:"ErrCode"`
	ErrInfo                string `schema:"ErrInfo"`
	Brand                  string `schema:"Brand"`
	DomesticFlag           string `schema:"DomesticFlag"`
	IssuerCode             string `schema:"IssuerCode"`
	DebitPrepaidFlag       string `schema:"DebitPrepaidFlag"`
	DebitPrepaidIssuerName string `schema:"DebitPrepaidIssuerName"`
	ForwardFinal           string `schema:"ForwardFinal"`
}

// SaveCard ... save card
func (cli *Client) SaveCard(
	req *SaveCardRequest,
) (*SaveCardResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &SaveCardResponse{}
	_, err := cli.do(saveCardPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteCardRequest ... delete card request
type DeleteCardRequest struct {
	MemberID string `schema:"MemberID" validate:"required,max=60"`
	SeqMode  string `schema:"SeqMode"`
	CardSeq  string `schema:"CardSeq"`
}

// Validate ... validate
func (r *DeleteCardRequest) Validate() error {
	return validate.Struct(r)
}

// DeleteCardResponse ... delete card response
type DeleteCardResponse struct {
	CardSeq string `schema:"CardSeq"`
	ErrCode string `schema:"ErrCode"`
	ErrInfo string `schema:"ErrInfo"`
}

// DeleteCard ... delete card
func (cli *Client) DeleteCard(
	req *DeleteCardRequest,
) (*DeleteCardResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := &DeleteCardResponse{}
	_, err := cli.do(deleteCardPath, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SearchCardRequest ...
type SearchCardRequest struct {
	MemberID        string `schema:"MemberID" validate:"required,max=60"`
	SeqMode         string `schema:"SeqMode"`
	CardSeq         string `schema:"CardSeq"`
	UseFloatingMask string `schema:"UseFloatingMask"`
}

// Validate ... validate
func (r *SearchCardRequest) Validate() error {
	return validate.Struct(r)
}

// SearchCardResponseDetail ... detail
type SearchCardResponseDetail struct {
	CardSeq                string `schema:"CardSeq"`
	DefaultFlag            string `schema:"DefaultFlag"`
	CardName               string `schema:"CardName"`
	CardNo                 string `schema:"CardNo"`
	Expire                 string `schema:"Expire"`
	HolderName             string `schema:"HolderName"`
	DeleteFlag             string `schema:"DeleteFlag"`
	ErrCode                string `schema:"ErrCode"`
	ErrInfo                string `schema:"ErrInfo"`
	Brand                  string `schema:"Brand"`
	DomesticFlag           string `schema:"DomesticFlag"`
	IssuerCode             string `schema:"IssuerCode"`
	DebitPrepaidFlag       string `schema:"DebitPrepaidFlag"`
	DebitPrepaidIssuerName string `schema:"DebitPrepaidIssuerName"`
	ForwardFinal           string `schema:"ForwardFinal"`
}

// SearchCardResponse ... search card response
type SearchCardResponse struct {
	Cards   []*SearchCardResponseDetail
	ErrCode string
	ErrInfo string
}

// SearchCard ... search card
func (cli *Client) SearchCard(
	req *SearchCardRequest,
) (*SearchCardResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	res := SearchCardResponseDetail{}
	_, err := cli.do(searchCardPath, req, &res)
	if err != nil {
		return nil, err
	}

	parsedResList := parser.ParseToMultiObject(res)

	var convertRes *SearchCardResponse
	convertRes.Cards = make([]*SearchCardResponseDetail, len(parsedResList))
	for _, parsedRes := range parsedResList {
		var dst SearchCardResponseDetail
		err = parser.MapToStruct(parsedRes, &dst)
		if err != nil {
			return nil, err
		}
		convertRes.Cards = append(convertRes.Cards, &dst)
	}

	convertRes.ErrCode = res.ErrCode
	convertRes.ErrInfo = res.ErrInfo

	return convertRes, nil
}
