package main

import (
	"github.com/apex/go-apex"
	// "github.com/kr/pretty"

	"encoding/json"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"

	"net/http"

	// "fmt"
	"log"
	// "os"
    "time"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {

		/* gather environment - Input Format of a Lambda Function for Proxy Integration */
		// http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-set-up-simple-proxy.html#api-gateway-simple-proxy-for-lambda-input-format

		var (
			resource              string
			path                  string
			httpMethod            string
			headers               map[string]interface{}
			queryStringParameters map[string]interface{}
			pathParameters        map[string]interface{}
			stageVariables        map[string]interface{}
			requestContext        map[string]interface{}
			body                  string
			isBase64Encoded       bool

			stagename string

			ok bool
		)

		var err error

		var t map[string]interface{}
		err = json.Unmarshal(event, &t)
		if err != nil {
			log.Println("parameters not provided: " + err.Error())
			return nil, err
		}

		log.Println(t)

		if resource, ok = t["resource"].(string); ok {
			log.Println("resource:", resource)
		} else {
			log.Println("Could not find resource")
		}

		if path, ok = t["path"].(string); ok {
			log.Println("path:", path)
		} else {
			log.Println("Could not find path")
		}

		if httpMethod, ok = t["httpMethod"].(string); ok {
			log.Println("httpMethod:", httpMethod)
		} else {
			log.Println("Could not find httpMethod")
		}

		if headers, ok = t["headers"].(map[string]interface{}); ok {
			log.Println("headers:", headers)
		} else {
			log.Println("Could not find headers")
		}

		if queryStringParameters, ok = t["queryStringParameters"].(map[string]interface{}); ok {
			log.Println("queryStringParameters:", queryStringParameters)
		} else {
			log.Println("Could not find queryStringParameters")
		}

		if pathParameters, ok = t["pathParameters"].(map[string]interface{}); ok {
			log.Println("pathParameters:", pathParameters)
		} else {
			log.Println("Could not find pathParameters")
		}

		if stageVariables, ok = t["stageVariables"].(map[string]interface{}); ok {
			log.Println("stageVariables:", stageVariables)

			stagename = stageVariables["stagename"].(string)
			log.Println("stagename:", stagename)

		} else {
			log.Println("Could not find stageVariables")
		}

		if requestContext, ok = t["requestContext"].(map[string]interface{}); ok {
			log.Println("requestContext:", requestContext)
		} else {
			log.Println("Could not find requestContext")
		}

		if body, ok = t["httpMethod"].(string); ok {
			log.Println("body:", body)
		} else {
			log.Println("Could not find body")
		}

		if isBase64Encoded, ok = t["isBase64Encoded"].(bool); ok {
			log.Println("isBase64Encoded:", isBase64Encoded)
		} else {
			log.Println("Could not find isBase64Encoded")
		}

		/* mostly mongodb */

		mapContext := make(map[string]interface{})

		var mongoSession *mgo.Session
		var mongoDatabase string

		mongoDatabase = "db_dev"
		if stagename == "prod" {
			mongoDatabase = "db_prod"
		}

		mapContext["mongoDatabase"] = mongoDatabase

		var mongoDialInfo *mgo.DialInfo

		mongoDialInfo = &mgo.DialInfo{
			// Addrs: []string{"<publicip>:27017"}, // you have to make sure the security groups have this port open
			Addrs: []string{"<privateip>:27017"}, // you have to configure your vpc config in lambda
			Timeout:  600 * time.Second,
			Database: "<authdb>",
			Username: "<user>",
			Password: "<password>",
		}

		mongoSession, err = mgo.DialWithInfo(mongoDialInfo)
		if err != nil {
			log.Fatal(err)
		}

		mongoSession.Close()

		/* Output Format of a Lambda Function for Proxy Integration */
		// http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-set-up-simple-proxy.html#api-gateway-simple-proxy-for-lambda-output-format

		return map[string]interface{}{"isBase64Encoded": false, "statusCode": http.StatusOK, "headers": map[string]interface{}{"headerName": "headerValue"}, "body": "sample_go"}, nil

	})
}

/*
map[

requestContext:map[httpMethod:GET path:/dev/tests/sample_apex_go accountId:126224339203 stage:dev identity:map[user:<nil> accountId:<nil> cognitoIdentityId:<nil> caller:<nil> sourceIp:112.198.72.203 accessKey:<nil> cognitoAuthenticationType:<nil> userAgent:curl/7.47.0 cognitoIdentityPoolId:<nil> apiKey: cognitoAuthenticationProvider:<nil> userArn:<nil>] resourceId:s6htmk requestId:16713c3f-2efe-11e7-8451-53e2dcbbcf04 resourcePath:/tests/sample_apex_go apiId:4ys084i6xb]


body:<nil>

isBase64Encoded:false

resource:/tests/sample_apex_go

headers:map[Accept:* /* CloudFront-Forwarded-Proto:https CloudFront-Is-Mobile-Viewer:false Via:1.1 13ce65b29463964b57d2dca8c7952968.cloudfront.net (CloudFront) X-Amz-Cf-Id:Ppso4YJlwD3LDBA3CvgJFNegHO-2s1-ji7ROF78WF_OYYt5US2me6Q== X-Forwarded-Proto:https CloudFront-Is-Desktop-Viewer:true CloudFront-Is-SmartTV-Viewer:false CloudFront-Is-Tablet-Viewer:false User-Agent:curl/7.47.0 X-Amzn-Trace-Id:Root=1-590822e9-2db685722d671d9237206492 X-Forwarded-Port:443 CloudFront-Viewer-Country:PH Host:4ys084i6xb.execute-api.ap-southeast-1.amazonaws.com X-Forwarded-For:112.198.72.203, 54.239.178.91]

queryStringParameters:<nil>
pathParameters:<nil>
stageVariables:map[stagename:dev database:db_dev]

path:/tests/sample_apex_go

httpMethod:GET

]

*/
