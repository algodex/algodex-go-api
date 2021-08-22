package design

import (
	"net/http"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API(
	"algodexidx", func() {
		Title("AlgoDex Indexer Service")
		Description("Service for tracking Algorand accounts and currently opted-in Holdings")
		cors.Origin("http://localhost")
		cors.Origin("http://algodex-go-api")
		cors.Origin(
			"/http[s]?://(.+[.])?algodex.com$/", func() {
				cors.Headers("*") //"X-Authorization", "X-Time", "X-Api-Version",
				//"Content-Type", "Origin",
				//"Authorization",

				cors.Methods("GET", "DELETE", "POST", "OPTIONS")
				//cors.Expose("Content-Type", "Origin")
				cors.MaxAge(600)
				//cors.Credentials()
			},
		)
		Server(
			"algodexidxsvr", func() {
				Host(
					"localhost", func() {
						URI("http://localhost:80")
					},
				)
			},
		)
	},
)

var Account = Type(
	"Account", func() {
		Description("Account describes an Algorand Account")
		Attribute(
			"address", String, "Public Account address", func() {
				MinLength(58)
				MaxLength(58)
				Pattern("^[A-Z2-7]{58}$")
				Example("4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU")
			},
		)
		Attribute("round", UInt64, "Round fetched")
		Attribute(
			"holdings", MapOf(String, Holding), func() {
				Description("Account Assets")
			},
		)
		Required("address", "round", "holdings")
	},
)

var Holding = Type(
	"Holding", func() {
		Description("Holding defines an ASA Asset ID and its balance.  ID 1 represents ALGO")
		Attribute(
			"asset", UInt64, func() {
				Description("ASA ID (1 for ALGO)")
				Example(uint64(202586210))
			},
		)
		Attribute(
			"amount", UInt64, func() {
				Description("Balance in asset base units")
			},
		)
		Attribute("decimals", UInt64)
		Attribute("metadataHash", String)
		Attribute("name", String, func() { Example("UNIT") })
		Attribute("unitName", String, func() { Example("My Unit") })
		Attribute("url", String, func() { Example("https://someurl.com") })
		Required("asset", "amount", "decimals", "metadataHash", "name", "unitName", "url")
	},
)

var TrackedAccount = ResultType(
	"application/vnd.algodex.account", func() {
		Description("A TrackedAccount is an Account returned by the indexer")
		Reference(Account)
		TypeName("TrackedAccount")
		Attributes(
			func() {
				Attribute("address")
				Attribute("round")
				Attribute("holdings")
			},
		)
		View(
			"default", func() {
				Attribute("address")
			},
		)
		View(
			"full", func() {
				Attribute("address")
				Attribute("round")
				Attribute("holdings")
			},
		)
	},
)

var addressList = ArrayOf(
	String, func() {
		MinLength(58)
		MaxLength(58)
		Pattern("^[A-Z2-7]{58}$")
		Example("4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU")
	},
)

var _ = Service(
	"account", func() {
		Description("The account service specifies which Algorand accounts to track")
		Error("access_denied")
		HTTP(
			func() {
				Response("access_denied", http.StatusUnauthorized)
			},
		)

		Method(
			"add", func() {
				Description("Add Algorand account(s) to track")
				Payload(
					func() {
						Attribute("address", addressList)
						Required("address")
					},
				)

				HTTP(
					func() {
						POST("/account")
						Response(StatusOK)
					},
				)
			},
		)
		Method(
			"delete", func() {
				Description("Delete Algorand account(s) to track")
				Payload(
					func() {
						Attribute("address", addressList)
						Required("address")
					},
				)

				HTTP(
					func() {
						DELETE("/account/{address}")
						Response(StatusOK)
					},
				)
			},
		)
		Method(
			"deleteAll", func() {
				Description("Delete all tracked algorand account(s).  Used for resetting everything")

				HTTP(
					func() {
						DELETE("/account/all")
						Response(StatusOK)
					},
				)
			},
		)

		Method(
			"get", func() {
				Description("Get specific account")
				Payload(
					func() {
						Attribute(
							"address", String, "Public Account address", func() {
								MinLength(58)
								MaxLength(58)
								Pattern("^[A-Z2-7]{58}$")
								Example("4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU")
							},
						)
						Required("address")
					},
				)
				Result(Account)
				HTTP(
					func() {
						GET("/account/{address}")
						Response(StatusOK)
					},
				)
			},
		)
		Method(
			"getMultiple", func() {
				Description("Get account(s)")
				Payload(
					func() {
						Attribute("address", addressList)
						Required("address")
					},
				)
				Result(ArrayOf(Account))
				HTTP(
					func() {
						POST("/account/get")
						Response(StatusOK)
					},
				)
			},
		)

		Method(
			"list", func() {
				Description("List all tracked accounts")
				Payload(
					func() {
						Attribute(
							"view", String, "View to render", func() {
								Enum("default", "full")
							},
						)
					},
				)
				Result(CollectionOf(TrackedAccount))
				HTTP(
					func() {
						GET("/account")
						Param("view")
						Response(StatusOK)
					},
				)
			},
		)

		Method(
			"isWatched", func() {
				Description("Returns which of the passed accounts are currently being monitored")
				Payload(
					func() {
						Attribute("address", addressList)
						Required("address")
					},
				)
				Result(ArrayOf(String))
				HTTP(
					func() {
						POST("/account/isWatched")
						Response(StatusOK)
					},
				)
			},
		)
	},
)

var Order = Type(
	"Order", func() {
		Description("Order is an individual buy or sell order")
		Attribute(
			"assetLimitPriceInAlgos", String,
			func() {
				Example(".08")
				Meta("struct:tag:db", "assetLimitPriceInAlgos")
				Meta("struct:tag:json", "assetLimitPriceInAlgos")
			},
		)
		Attribute(
			"asaPrice", String, func() {
				Example(".08")
				Meta("struct:tag:db", "asaPrice")
				Meta("struct:tag:json", "asaPrice")
			},
		)
		Attribute(
			"assetLimitPriceD", UInt64, func() {
				Example(197)
				Meta("struct:tag:db", "assetLimitPriceD")
				Meta("struct:tag:json", "assetLimitPriceD")
			},
		)
		Attribute(
			"assetLimitPriceN", UInt64, func() {
				Example(100)
				Meta("struct:tag:db", "assetLimitPriceN")
				Meta("struct:tag:json", "assetLimitPriceN")
			},
		)
		Attribute(
			"algoAmount", UInt64, func() {
				Example(498000)
				Meta("struct:tag:db", "algoAmount")
				Meta("struct:tag:json", "algoAmount")
			},
		)
		Attribute(
			"asaAmount", UInt64, func() {
				Example(1000000)
				Meta("struct:tag:db", "asaAmount")
				Meta("struct:tag:json", "asaAmount")
			},
		)
		Attribute(
			"assetId", UInt64, func() {
				Example(15322902)
				Meta("struct:tag:db", "assetId")
				Meta("struct:tag:json", "assetId")
			},
		)
		Attribute(
			"appId", UInt64, func() {
				Example(16021157)
				Meta("struct:tag:db", "appId")
				Meta("struct:tag:json", "appId")
			},
		)
		Attribute(
			"escrowAddress", String, func() {
				Example("2IYBUR4WXPWGBKRETN4GVSCPG7VOJRMVZFYTDYMQSRMXQJY24EHGFLFIMU")
				Meta("struct:tag:db", "escrowAddress")
				Meta("struct:tag:json", "escrowAddress")
			},
		)
		Attribute(
			"ownerAddress", String, func() {
				Example("XHGANA4SOVZKH4GGSSLAMOZDVWWVIXT5DZBIEGI3GX2EESVFNFGFTHJATA")
				Meta("struct:tag:db", "ownerAddress")
				Meta("struct:tag:json", "ownerAddress")
			},
		)
		Attribute(
			"minimumExecutionSizeInAlgo", UInt64, func() {
				Example(0)
				Meta("struct:tag:db", "minimumExecutionSizeInAlgo")
				Meta("struct:tag:json", "minimumExecutionSizeInAlgo")
			},
		)
		Attribute(
			"round", UInt64, func() {
				Example(16043694)
				Meta("struct:tag:db", "round")
				Meta("struct:tag:json", "round")
			},
		)
		Attribute(
			"unix_time", UInt64, func() {
				Example(1629064223)
				Meta("struct:tag:db", "unix_time")
				Meta("struct:tag:json", "unix_time")
			},
		)
		Attribute(
			"formattedPrice", String, func() {
				Example("1.970000")
				Meta("struct:tag:db", "formattedPrice")
				Meta("struct:tag:json", "formattedPrice")
			},
		)
		Attribute(
			"formattedASAAmount", String, func() {
				Example("1.000000")
				Meta("struct:tag:db", "formattedASAAmount")
				Meta("struct:tag:json", "formattedASAAmount")
			},
		)
		Attribute(
			"decimals", UInt64, func() {
				Example(6)
				Meta("struct:tag:db", "decimals")
				Meta("struct:tag:json", "decimals")
			},
		)
		Required(
			"assetLimitPriceInAlgos",
			"asaPrice",
			"assetLimitPriceD",
			"assetLimitPriceN",
			"algoAmount",
			"asaAmount",
			"assetId",
			"appId",
			"escrowAddress",
			"ownerAddress",
			"minimumExecutionSizeInAlgo",
			"round",
			"unix_time",
			"formattedPrice",
			"formattedASAAmount",
			"decimals",
		)
	},
)

var Orders = Type(
	"Orders", func() {
		Description("Orders contains a list of buy/sell orders matching the criteria.")
		Attribute("sellASAOrdersInEscrow", ArrayOf(Order), func() { Description("Sell orders") })
		Attribute("buyASAOrdersInEscrow", ArrayOf(Order), func() { Description("Buy orders") })
	},
)

var _ = Service(
	"orders", func() {
		Description("The orders service provides information on open orders")
		Error("access_denied")
		Error("missing_parameters")
		HTTP(
			func() {
				Response("access_denied", http.StatusUnauthorized)
			},
		)

		Method(
			"get", func() {
				Description("Get all open orders for a specific asset")
				Payload(
					func() {
						Attribute(
							"assetId", UInt64, "ASA ID", func() {
								Example(15322902)
							},
						)
						Attribute("ownerAddr", addressList, "Owner address(es)")
					},
				)
				Result(Orders)
				HTTP(
					func() {
						GET("/orders")
						Param("assetId")
						Param("ownerAddr")
						Response(StatusOK)
					},
				)
			},
		)
	},
)
var _ = Service(
	"inspect", func() {
		Description("The inspect service provides msgpack decoding services")
		Error("access_denied")
		HTTP(
			func() {
				Response("access_denied", http.StatusUnauthorized)
			},
		)

		Method(
			"unpack", func() {
				Description("Unpack a msgpack body (base64 encoded) returning 'goal clerk inspect' output")
				Payload(
					func() {
						Attribute("msgpack", String)
					},
				)
				Result(
					String, func() {
						Description("Returns output from goal clerk inspect of passed msgpack-encoded payload")
					},
				)

				HTTP(
					func() {
						POST("/inspect/unpack")
						Response(
							StatusOK, func() {
								ContentType("text/plain")
							},
						)
					},
				)
			},
		)
	},
)

var _ = Service(
	"info", func() {
		Description("The info service provides information on version data, liveness, readiness checks, etc.")

		Method(
			"version", func() {
				Description("Returns version information for the service")
				Result(
					String, func() {
						Example("14193a3-dirty")
					},
				)
				HTTP(
					func() {
						GET("/version")
						Response(StatusOK)
					},
				)
			},
		)
		Method(
			"live", func() {
				Description("Simple health check")
				HTTP(
					func() {
						GET("/live")
						Response(StatusOK)
					},
				)
			},
		)
		Files("/openapi3.yaml", "./openapi3.yaml")
	},
)
