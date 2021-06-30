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
			"assets", ArrayOf(Int64), func() {
				Description("Opted-in ASA IDs")
				Example([]int64{202586210, 205471981})
			},
		)
		//Field(2, "assets", ArrayOf(Int64), "List of opted-in Assets")
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
				//Attribute("assets")
				//	Attribute("assets", func() {
				//		View("full")
				//	})
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

//var TrackedAccount = ResultType("application/vnd.cellar.stored-bottle", func() {
//	Description("A StoredBottle describes a bottle retrieved by the storage service.")
//	Reference(Bottle)
//	TypeName("StoredBottle")
//
//	Attributes(func() {
//		Attribute("id", String, "ID is the unique id of the bottle.", func() {
//			Example("123abc")
//			Meta("rpc:tag", "8")
//		})
//		Field(2, "name")
//		Field(3, "winery")
//		Field(4, "vintage")
//		Field(5, "composition")
//		Field(6, "description")
//		Field(7, "rating")
//	})
//
//	View("default", func() {
//		Attribute("id")
//		Attribute("name")
//		Attribute("winery", func() {
//			View("tiny")
//		})
//		Attribute("vintage")
//		Attribute("composition")
//		Attribute("description")
//		Attribute("rating")
//	})
//
//	View("tiny", func() {
//		Attribute("id")
//		Attribute("name")
//		Attribute("winery", func() {
//			View("tiny")
//		})
//	})
//
//	Required("id", "name", "winery", "vintage")
//})

var _ = Service(
	"account", func() {
		Description("The account service specifies which Algorand accounts to track")

		Method(
			"add", func() {
				Description("Add Algorand account to track")
				Payload(String)

				//Result(String)
				Result(Account)

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
				//Result(ArrayOf(Account))
				Result(CollectionOf(TrackedAccount))
				HTTP(
					func() {
						GET("/")
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
