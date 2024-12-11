package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"

	"github.com/chaihaobo/go-graphql-template/resolver"
	"github.com/chaihaobo/go-graphql-template/schema"
)

func main() {
	// Tweak configuration values here.
	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	rootResolver := resolver.New()
	s, err := schema.String()
	if err != nil {
		log.Fatalf("reading embedded schema contents: %s", err)
	}
	schema := graphql.MustParseSchema(s, rootResolver)
	// Register handlers to routes.
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	graphQLHandler := graphqlws.NewHandlerFunc(schema, &relay.Handler{Schema: schema})
	mux.Handle("/graphql/", graphQLHandler)
	mux.Handle("/graphql", graphQLHandler) // Register without a trailing slash to avoid redirect.

	// Configure the HTTP server.
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	// Begin listeing for requests.
	log.Printf("Listening for requests on %s", srv.Addr)

	if err = srv.ListenAndServe(); err != nil {
		log.Println("server.ListenAndServe:", err)
	}

	// TODO: intercept shutdown signals for cleanup of connections.
	log.Println("Shut down.")

}

var page = []byte(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>GraphiQL</title>
    <style>
      body {
        height: 100%;
        margin: 0;
        width: 100%;
        overflow: hidden;
      }
      #graphiql {
        height: 100vh;
      }
    </style>
    <script src="https://unpkg.com/react@17/umd/react.development.js" integrity="sha512-Vf2xGDzpqUOEIKO+X2rgTLWPY+65++WPwCHkX2nFMu9IcstumPsf/uKKRd5prX3wOu8Q0GBylRpsDB26R6ExOg==" crossorigin="anonymous"></script>
    <script src="https://unpkg.com/react-dom@17/umd/react-dom.development.js" integrity="sha512-Wr9OKCTtq1anK0hq5bY3X/AvDI5EflDSAh0mE9gma+4hl+kXdTJPKZ3TwLMBcrgUeoY0s3dq9JjhCQc7vddtFg==" crossorigin="anonymous"></script>
    <link rel="stylesheet" href="https://unpkg.com/graphiql@2.3.0/graphiql.min.css" />
  </head>
  <body>
    <div id="graphiql">Loading...</div>
    <script src="https://unpkg.com/graphiql@2.3.0/graphiql.min.js" type="application/javascript"></script>
	<script src="https://cdn.jsdelivr.net/npm/subscriptions-transport-ws@0.11.0/browser/client.js"></script>
    <script>
	const wsClient = new window.SubscriptionsTransportWs.SubscriptionClient(
      'ws://localhost:8000/graphql',
      { reconnect: true }
    );
      ReactDOM.render(
        React.createElement(GraphiQL, {
          //fetcher: GraphiQL.createFetcher({url: '/graphql',subscriptionUrl:'ws://localhost:8000/graphql', protocol: "graphql-ws",legacyClient:wsClient}),
          fetcher: GraphiQL.createFetcher({url: '/graphql',legacyClient:wsClient}),
          defaultEditorToolsVisibility: true,
        }),
        document.getElementById('graphiql'),
      );
    </script>
  </body>
</html>
`)
