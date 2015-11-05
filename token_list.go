/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Token"}
  on: Thu Nov 05 10:16:52 +0700 2015
  by: chakrit
*/

package omise

// TokenList represents the list structure returned by Omise's REST API that contains
// Token struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TokenList struct {
	List
  Data []*Token `json:"data"`
}

// Find finds and returns Token with the given id. Returns nil if not found.
func (list *TokenList) Find(id string) *Token {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
