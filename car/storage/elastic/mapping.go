package elasticstore

func mapping(indexName string) string {

	return `{
	"settings": {
		"analysis": {
			"analyzer": {
				"my_analyzer": {
					"type": "custom",
					"filter": [
						"lowercase"
					],
					"tokenizer": "whitespace"
				}
			}
		}
	},
	"mappings":{
		"properties":{
			"Make":{
				"type": "text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Model":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Year":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Color":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Class":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Available":{
				"type":"bool",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"AccidentHistory": {
				"type":"nested", 
				"properties":{
					"Date":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"Details":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
				}	
			},
			
			
		}
	}	
}
`
}
