package management

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

var logStreamTestCases = []logStreamTestCase{
	{
		name: "AmazonEventBridge LogStream",
		logStream: LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeAmazonEventBridge),
			Sink: &LogStreamSinkAmazonEventBridge{
				AccountID: auth0.String("999999999999"),
				Region:    auth0.String("us-west-2"),
			},
		},
	},
	// { // This test requires an active subscription.
	// 	name: "AzureEventGrid LogStream",
	// 	logStream: LogStream{
	// 		Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
	// 		Type: auth0.String(LogStreamTypeAzureEventGrid),
	// 		Sink: &LogStreamSinkAzureEventGrid{
	// 			SubscriptionID: auth0.String("b69a6835-57c7-4d53-b0d5-1c6ae580b6d5"),
	// 			Region:         auth0.String("northeurope"),
	// 			ResourceGroup:  auth0.String("azure-logs-rg"),
	// 		},
	// 	},
	// },
	{
		name: "HTTP LogStream",
		logStream: LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeHTTP),
			Sink: &LogStreamSinkHTTP{
				Endpoint:      auth0.String("https://example.com/logs"),
				Authorization: auth0.String("Bearer f2368bbe77074527a37be2fdd5b92bad"),
				ContentFormat: auth0.String("JSONLINES"),
				ContentType:   auth0.String("application/json"),
			},
		},
	},
	{
		name: "DataDog LogStream",
		logStream: LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeDatadog),
			Sink: &LogStreamSinkDatadog{
				APIKey: auth0.String("121233123455"),
				Region: auth0.String("us"),
			},
		},
	},
	{
		name: "Splunk LogStream",
		logStream: LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeSplunk),
			Sink: &LogStreamSinkSplunk{
				Domain: auth0.String("demo.splunk.com"),
				Port:   auth0.String("8080"),
				Secure: auth0.Bool(true),
				Token:  auth0.String("12a34ab5-c6d7-8901-23ef-456b7c89d0c1"),
			},
		},
	},
	{
		name: "Sumo LogStream",
		logStream: LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeSumo),
			Sink: &LogStreamSinkSumo{
				SourceAddress: auth0.String("https://example.com"),
			},
		},
	},
}

type logStreamTestCase struct {
	name      string
	logStream LogStream
}

func TestLogStreamManager_Create(t *testing.T) {
	for _, testCase := range logStreamTestCases {
		t.Run("It can successfully create a "+testCase.name, func(t *testing.T) {
			setupHTTPRecordings(t)

			expectedLogStream := testCase.logStream

			err := m.LogStream.Create(&expectedLogStream)
			assert.NoError(t, err)
			assert.NotEmpty(t, expectedLogStream.GetID())

			t.Cleanup(func() {
				cleanupLogStream(t, expectedLogStream.GetID())
			})
		})
	}
}

func TestLogStreamManager_Read(t *testing.T) {
	for _, testCase := range logStreamTestCases {
		t.Run("It can successfully read a "+testCase.name, func(t *testing.T) {
			setupHTTPRecordings(t)

			expectedLogStream := givenALogStream(t, testCase)

			actualLogStream, err := m.LogStream.Read(expectedLogStream.GetID())

			assert.NoError(t, err)
			assert.Equal(t, expectedLogStream, actualLogStream)
		})
	}
}

func TestLogStreamManager_Update(t *testing.T) {
	for _, testCase := range logStreamTestCases {
		t.Run("It can successfully update a "+testCase.name, func(t *testing.T) {
			setupHTTPRecordings(t)

			logStream := givenALogStream(t, testCase)
			updatedLogStream := &LogStream{
				Filters: []interface{}{
					map[string]string{
						"type": "category",
						"name": "auth.login.fail",
					},
				},
			}

			err := m.LogStream.Update(logStream.GetID(), updatedLogStream)
			assert.NoError(t, err)

			actualLogStream, err := m.LogStream.Read(logStream.GetID())
			assert.NoError(t, err)
			assert.Equal(t, updatedLogStream.Filters, actualLogStream.Filters)
		})
	}
}

func TestLogStreamManager_Delete(t *testing.T) {
	for _, testCase := range logStreamTestCases {
		t.Run("It can successfully delete a "+testCase.name, func(t *testing.T) {
			setupHTTPRecordings(t)

			logStream := givenALogStream(t, testCase)

			err := m.LogStream.Delete(logStream.GetID())
			assert.NoError(t, err)

			actualLogStream, err := m.LogStream.Read(logStream.GetID())
			assert.Nil(t, actualLogStream)
			assert.Error(t, err)
			assert.Implements(t, (*Error)(nil), err)
			assert.Equal(t, http.StatusNotFound, err.(Error).Status())
		})
	}
}

func TestLogStreamManager_List(t *testing.T) {
	setupHTTPRecordings(t)

	// There are no params we can add here, unfortunately.
	logStreamList, err := m.LogStream.List()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(logStreamList), 0)
}

func givenALogStream(t *testing.T, testCase logStreamTestCase) *LogStream {
	t.Helper()

	logStream := testCase.logStream

	err := m.LogStream.Create(&logStream)
	require.NoError(t, err)

	t.Cleanup(func() {
		cleanupLogStream(t, logStream.GetID())
	})

	return &logStream
}

func cleanupLogStream(t *testing.T, logStreamID string) {
	t.Helper()

	err := m.LogStream.Delete(logStreamID)
	require.NoError(t, err)
}
