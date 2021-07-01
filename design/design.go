package design

import . "goa.design/goa/v3/dsl"

var _ = API(
	"algodexidx", func() {
		Title("AlgoDex Indexer Service")
		Description("Service for tracking Algorand accounts and currently opted-in Assets")
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
			"assets", ArrayOf(UInt64), func() {
				Description("Opted-in ASA IDs")
				Example([]uint64{202586210, 205471981})
			},
		)
		Required("address", "assets")
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
				Attribute("assets")
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
				Attribute("assets")
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

				//Result(String)
				//Result(Account)

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
		//var _ = Service("openapi", func() {
		//	Files("/swagger.json", "./gen/http/openapi.json")
		Files("/openapi.json", "./gen/http/openapi.json")
		//})
	},
)
