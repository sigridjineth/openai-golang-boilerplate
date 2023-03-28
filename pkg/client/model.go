package client

// ModelObject contained in an client reponse
/*
curl https://api.openai.com/v1/models \
  -H 'Authorization: Bearer YOUR_API_KEY'

{
    "object": "list",
    "data": [
        {
            "id": "babbage",
            "object": "client",
            "created": 1649358449,
            "owned_by": "openai",
            "permission": [
                {
                    "id": "modelperm-49FUp5v084tBB49tC4z8LPH5",
                    "object": "model_permission",
                    "created": 1669085501,
                    "allow_create_engine": false,
                    "allow_sampling": true,
                    "allow_logprobs": true,
                    "allow_search_indices": false,
                    "allow_view": true,
                    "allow_fine_tuning": false,
                    "organization": "*",
                    "group": null,
                    "is_blocking": false
                }
            ],
            "root": "babbage",
            "parent": null
        },
        {
            "id": "ada",
            "object": "client",
            "created": 1649357491,
            "owned_by": "openai",
            "permission": [
                {
                    "id": "modelperm-xTOEYvDZGN7UDnQ65VpzRRHz",
                    "object": "model_permission",
                    "created": 1669087301,
                    "allow_create_engine": false,
                    "allow_sampling": true,
                    "allow_logprobs": true,
                    "allow_search_indices": false,
                    "allow_view": true,
                    "allow_fine_tuning": false,
                    "organization": "*",
                    "group": null,
                    "is_blocking": false
                }
            ],
            "root": "ada",
            "parent": null
        },
        {
            "id": "davinci",
            "object": "client",
            "created": 1649359874,
            "owned_by": "openai",
            "permission": [
                {
                    "id": "modelperm-U6ZwlyAd0LyMk4rcMdz33Yc3",
                    "object": "model_permission",
                    "created": 1669066355,
                    "allow_create_engine": false,
                    "allow_sampling": true,
                    "allow_logprobs": true,
                    "allow_search_indices": false,
                    "allow_view": true,
                    "allow_fine_tuning": false,
                    "organization": "*",
                    "group": null,
                    "is_blocking": false
                }
            ],
            "root": "davinci",
            "parent": null
        },

*/

const ModelEndPoint = "/models"
const GetAllModels = ModelEndPoint + "/all"
const RetrieveModels = ModelEndPoint + "/retrieve"
const ModelIdParamKey = "model_id"

type PermissionInModelObject struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Created            int         `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}

type ModelObject struct {
	ID         string                    `json:"id"`
	Object     string                    `json:"object"`
	OwnedBy    string                    `json:"owned_by"`
	Permission []PermissionInModelObject `json:"permission"`
	Root       string                    `json:"root"`
	Parent     interface{}               `json:"parent"`
}

type ListModelsResponse struct {
	Data   []ModelObject `json:"data"`
	Object string        `json:"object"`
}
