package prerender

var (
	// serviceURL is the default URL of Prerender service
	serviceURL = "https://service.prerender.io/"

	// CrawlerUserAgents are list of bot UAs
	CrawlerUserAgents = []string{
		"googlebot",
		"Yahoo! Slurp",
		"bingbot",
		"yandex",
		"baiduspider",
		"facebookexternalhit",
		"twitterbot",
		"rogerbot",
		"linkedinbot",
		"embedly",
		"quora link preview",
		"showyoubot",
		"outbrain",
		"pinterest/0.",
		"developers.google.com/+/web/snippet",
		"slackbot",
		"vkShare",
		"W3C_Validator",
		"redditbot",
		"Applebot",
		"WhatsApp",
		"flipboard",
		"tumblr",
		"bitlybot",
		"SkypeUriPreview",
		"nuzzel",
		"Discordbot",
		"Google Page Speed",
		"Qwantify",
		"pinterestbot",
		"Bitrix link preview",
		"XING-contenttabreceiver",
		"Chrome-Lighthouse",
	}

	// ExtensionsToIgnore are file extensions that we won't send to Prerender
	ExtensionsToIgnore = []string{
		".js",
		".css",
		".xml",
		".less",
		".png",
		".jpg",
		".jpeg",
		".gif",
		".pdf",
		".doc",
		".txt",
		".ico",
		".rss",
		".zip",
		".mp3",
		".rar",
		".exe",
		".wmv",
		".doc",
		".avi",
		".ppt",
		".mpg",
		".mpeg",
		".tif",
		".wav",
		".mov",
		".psd",
		".ai",
		".xls",
		".mp4",
		".m4a",
		".swf",
		".dat",
		".dmg",
		".iso",
		".flv",
		".m4v",
		".torrent",
		".woff",
		".ttf",
		".svg",
		".webmanifest",
	}
)
