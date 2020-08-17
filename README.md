# xml2json as a service

## Convert online xml file (like rss) to pretty json response 

With this service you can call the endpoint for convert operation. Additionally you can save response as txt file.

App runs on : `localhost:1323`

---

Sample Endpoint

`POST` : `localhost:1323/xml`

Sample Request
```json
{
    "link":"http://feeds.bbci.co.uk/news/technology/rss.xml#",
    "write":true
}
```

Sample Response
```json
{
	"rss": {
		"-atom": "http://www.w3.org/2005/Atom",
		"-content": "http://purl.org/rss/1.0/modules/content/",
		"-dc": "http://purl.org/dc/elements/1.1/",
		"-media": "http://search.yahoo.com/mrss/",
		"-version": "2.0",
		"channel": {
			"copyright": "Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.",
			"description": "BBC News - Technology",
			"generator": "RSS for Node",
			"image": {
				"link": "https://www.bbc.co.uk/news/",
				"title": "BBC News - Technology",
				"url": "https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif"
			},
			"item": [
				//...
				{
					"description": "From insurance and healthcare to social media and policing, algorithms are used all around us.",
					"guid": {
						"#content": "https://www.bbc.co.uk/news/technology-53806038",
						"-isPermaLink": "true"
					},
					"link": "https://www.bbc.co.uk/news/technology-53806038",
					"pubDate": "Mon, 17 Aug 2020 14:12:01 GMT",
					"title": "The algorithms that make big decisions about your life"
				},
				{
					"description": "The tech firm is fighting efforts to make it pay for local media content.",
					"guid": {
						"#content": "https://www.bbc.co.uk/news/technology-53806489",
						"-isPermaLink": "true"
					},
					"link": "https://www.bbc.co.uk/news/technology-53806489",
					"pubDate": "Mon, 17 Aug 2020 11:23:38 GMT",
					"title": "Google says Australian law would put search and YouTube at risk"
				},
				{
					"description": "The actions will stop the firm buying computer chips made using US tech, even if made abroad.",
					"guid": {
						"#content": "https://www.bbc.co.uk/news/business-53805038",
						"-isPermaLink": "true"
					},
					"link": "https://www.bbc.co.uk/news/business-53805038",
					"pubDate": "Mon, 17 Aug 2020 15:15:46 GMT",
					"title": "Huawei: US tightens restrictions on Chinese giant"
				}
				//...
			],
			"language": "en-gb",
			"lastBuildDate": "Mon, 17 Aug 2020 21:35:13 GMT",
			"link": "https://www.bbc.co.uk/news/",
			"title": "BBC News - Technology",
			"ttl": "15"
		}
	}
}
```
