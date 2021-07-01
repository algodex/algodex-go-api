package design

import . "goa.design/goa/v3/dsl"

var _ = API(
	"algodexidx", func() {
		Title("AlgoDex Indexer Service")
		Description("Service for tracking Algorand accounts and currently opted-in Holdings")
		Server(
			"algodexidxsvr", func() {
				Host(
					"localhost", func() {
						URI("http://localhost:8000")
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
				MaxLength(58)
				Example("4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU")
			},
		)
		Attribute(
			"holdings", MapOf(String, UInt64), func() {
				Description("Opted-in ASA IDs")
				Example(map[string]uint64{"202586210": 100, "205471981": 200})
			},
		)
		Required("address", "holdings")
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
				Attribute("holdings")
			},
		)
	},
)

var _ = Service(
	"account", func() {
		Description("The account service specifies which Algorand accounts to track")

		Method(
			"add", func() {
				Description("Add Algorand account to track")
				Payload(String)

				HTTP(
					func() {
						POST("/")
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
								MaxLength(58)
								Example("4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU")
							},
						)
						Required("address")
					},
				)
				Result(Account)
				HTTP(
					func() {
						GET("/{address}")
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
						Field(
							1, "view", String, "View to render", func() {
								Enum("default", "full")
							},
						)
					},
				)
				Result(CollectionOf(TrackedAccount))
				HTTP(
					func() {
						GET("/")
						Param("view")
						Response(StatusOK)
					},
				)
			},
		)
		Files("/openapi.json", "./gen/http/openapi.json")
	},
)
