package management

import (
	"testing"
	"time"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/internal/testing/expect"
)

func TestLogStream(t *testing.T) {
	l := &LogStream{
		Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
		Type: auth0.String(LogStreamTypeDatadog),
		Sink: &LogStreamSinkDatadog{
			APIKey: auth0.String("12334567876543"),
			Region: auth0.String("eu"),
		},
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := l.Sink.(*LogStreamSinkDatadog); !ok {
			t.Errorf("unexpected options type %T", l.Sink)
		}
		t.Logf("%v\n", l)
	})

	t.Run("Read", func(t *testing.T) {
		l, err = m.LogStream.Read(l.GetID())
		if err != nil {
			t.Error(err)
		}
		if _, ok := l.Sink.(*LogStreamSinkDatadog); !ok {
			t.Errorf("unexpected options type %T", l.Sink)
		}
		t.Logf("%v\n", l)
	})

	t.Run("List", func(t *testing.T) {
		lsl, err := m.LogStream.List()
		if err != nil {
			t.Error(err)
		}
		for _, ls := range lsl {
			var ok bool

			switch ls.GetType() {
			case LogStreamTypeAmazonEventBridge:
				_, ok = ls.Sink.(*LogStreamSinkAmazonEventBridge)
			case LogStreamTypeAzureEventGrid:
				_, ok = ls.Sink.(*LogStreamSinkAzureEventGrid)
			case LogStreamTypeHTTP:
				_, ok = ls.Sink.(*LogStreamSinkHTTP)
			case LogStreamTypeDatadog:
				_, ok = ls.Sink.(*LogStreamSinkDatadog)
			case LogStreamTypeSplunk:
				_, ok = ls.Sink.(*LogStreamSinkSplunk)
			case LogStreamTypeSumo:
				_, ok = ls.Sink.(*LogStreamSinkSumo)
			default:
				_, ok = ls.Sink.(map[string]interface{})
			}

			if !ok {
				t.Errorf("unexpected options type %T", ls.Sink)
			}

			t.Logf("%s %s %T\n", ls.GetID(), ls.GetName(), ls.Sink)
		}
	})

	t.Run("Update", func(t *testing.T) {
		id := l.GetID()

		l.ID = nil   // read-only
		l.Name = nil // read-only
		l.Type = nil // read-only

		l.Sink = &LogStreamSinkDatadog{
			APIKey: auth0.String("12334567876543"),
			Region: auth0.String("us"),
		}

		err = m.LogStream.Update(id, l)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", l)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.LogStream.Delete(l.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestLogStreamSink(t *testing.T) {
	t.Run("AmazonEventBridge", func(t *testing.T) {
		l := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeAmazonEventBridge),
			Sink: &LogStreamSinkAmazonEventBridge{
				AccountID: auth0.String("999999999999"),
				Region:    auth0.String("us-west-2"),
			},
		}

		defer func() { m.LogStream.Delete(l.GetID()) }()

		err := m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}

		s, ok := l.Sink.(*LogStreamSinkAmazonEventBridge)
		if !ok {
			t.Fatalf("unexpected type %T", s)
		}

		expect.Expect(t, l.GetStatus(), "active")
		expect.Expect(t, l.GetType(), LogStreamTypeAmazonEventBridge)
		expect.Expect(t, s.GetAccountID(), "999999999999")
		expect.Expect(t, s.GetRegion(), "us-west-2")
		expect.Expect(t, len(s.GetPartnerEventSource()) > 0, true)

		t.Logf("%s\n", l)
	})

	t.Run("AzureEventGrid", func(t *testing.T) {
		t.Skip("this test requires an active subscription")

		l := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeAzureEventGrid),
			Sink: &LogStreamSinkAzureEventGrid{
				SubscriptionID: auth0.String("b69a6835-57c7-4d53-b0d5-1c6ae580b6d5"),
				Region:         auth0.String("northeurope"),
				ResourceGroup:  auth0.String("azure-logs-rg"),
			},
		}

		defer func() { m.LogStream.Delete(l.GetID()) }()

		err := m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}

		s, ok := l.Sink.(*LogStreamSinkAzureEventGrid)
		if !ok {
			t.Fatalf("unexpected type %T", s)
		}

		expect.Expect(t, s.GetSubscriptionID(), "b69a6835-57c7-4d53-b0d5-1c6ae580b6d5")
		expect.Expect(t, s.GetRegion(), "northeurope")
		expect.Expect(t, s.GetResourceGroup(), "azure-logs-rg")
		expect.Expect(t, len(s.GetPartnerTopic()) > 0, true)

		t.Logf("%s\n", l)
	})

	t.Run("HTTP", func(t *testing.T) {
		l := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeHTTP),
			Sink: &LogStreamSinkHTTP{
				Endpoint:      auth0.String("https://example.com/logs"),
				Authorization: auth0.String("Bearer f2368bbe77074527a37be2fdd5b92bad"),
				ContentFormat: auth0.String("JSONLINES"),
				ContentType:   auth0.String("application/json"),
			},
		}

		defer func() { m.LogStream.Delete(l.GetID()) }()

		err := m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}

		s, ok := l.Sink.(*LogStreamSinkHTTP)
		if !ok {
			t.Fatalf("unexpected type %T", s)
		}

		expect.Expect(t, l.GetStatus(), "active")
		expect.Expect(t, l.GetType(), LogStreamTypeHTTP)
		expect.Expect(t, s.GetEndpoint(), "https://example.com/logs")
		expect.Expect(t, s.GetAuthorization(), "Bearer f2368bbe77074527a37be2fdd5b92bad")
		expect.Expect(t, s.GetContentFormat(), "JSONLINES")
		expect.Expect(t, s.GetContentType(), "application/json")

		t.Logf("%s\n", l)
	})

	t.Run("Datadog", func(t *testing.T) {
		l := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeDatadog),
			Sink: &LogStreamSinkDatadog{
				APIKey: auth0.String("121233123455"),
				Region: auth0.String("us"),
			},
		}

		defer func() { m.LogStream.Delete(l.GetID()) }()

		err := m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}

		s, ok := l.Sink.(*LogStreamSinkDatadog)
		if !ok {
			t.Fatalf("unexpected type %T", s)
		}

		expect.Expect(t, l.GetStatus(), "active")
		expect.Expect(t, l.GetType(), LogStreamTypeDatadog)
		expect.Expect(t, s.GetAPIKey(), "121233123455")
		expect.Expect(t, s.GetRegion(), "us")

		t.Logf("%s\n", l)
	})

	t.Run("Splunk", func(t *testing.T) {
		l := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeSplunk),
			Sink: &LogStreamSinkSplunk{
				Domain: auth0.String("demo.splunk.com"),
				Port:   auth0.String("8080"),
				Secure: auth0.Bool(true),
				Token:  auth0.String("12a34ab5-c6d7-8901-23ef-456b7c89d0c1"),
			},
		}

		defer func() { m.LogStream.Delete(l.GetID()) }()

		err := m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}

		s, ok := l.Sink.(*LogStreamSinkSplunk)
		if !ok {
			t.Fatalf("unexpected type %T", s)
		}

		expect.Expect(t, l.GetStatus(), "active")
		expect.Expect(t, l.GetType(), LogStreamTypeSplunk)
		expect.Expect(t, s.GetDomain(), "demo.splunk.com")
		expect.Expect(t, s.GetPort(), "8080")
		expect.Expect(t, s.GetSecure(), true)
		expect.Expect(t, s.GetToken(), "12a34ab5-c6d7-8901-23ef-456b7c89d0c1")

		t.Logf("%s\n", l)
	})

	t.Run("Sumo", func(t *testing.T) {
		l := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamTypeSumo),
			Sink: &LogStreamSinkSumo{
				SourceAddress: auth0.String("https://example.com"),
			},
		}

		defer func() { m.LogStream.Delete(l.GetID()) }()

		err := m.LogStream.Create(l)
		if err != nil {
			t.Fatal(err)
		}

		s, ok := l.Sink.(*LogStreamSinkSumo)
		if !ok {
			t.Fatalf("unexpected type %T", s)
		}

		expect.Expect(t, l.GetStatus(), "active")
		expect.Expect(t, l.GetType(), LogStreamTypeSumo)
		expect.Expect(t, s.GetSourceAddress(), "https://example.com")

		t.Logf("%s\n", l)
	})
}
