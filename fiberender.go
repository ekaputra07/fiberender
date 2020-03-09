package fiberender

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber"
)

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

// PrerenderConfig can be used to customize Prerender
type PrerenderConfig struct {
	Token              string
	ServiceURL         string
	Host               string
	ForwardHeaders     bool
	Protocol           string
	CrawlerUserAgents  []string
	ExtensionsToIgnore []string
	Whitelist          []string
	Blacklist          []string

	// TODO:
	// BeforeRender() https://github.com/prerender/prerender-node#beforerender
	// AfterRender() https://github.com/prerender/prerender-node#afterrender
}

// defaultConfig is the default configuration if user not specify their own config
var defaultConfig = PrerenderConfig{
	ServiceURL:         serviceURL,
	CrawlerUserAgents:  CrawlerUserAgents,
	ExtensionsToIgnore: ExtensionsToIgnore,
}

// New returns the middleware
func New(config ...PrerenderConfig) func(*fiber.Ctx) {
	var cfg PrerenderConfig

	if len(config) == 0 {
		cfg = defaultConfig
	} else {
		cfg = config[0]
		// set default values if not set
		if cfg.ServiceURL == "" {
			cfg.ServiceURL = serviceURL
		}
		if len(cfg.CrawlerUserAgents) == 0 {
			cfg.CrawlerUserAgents = CrawlerUserAgents
		}
		if len(cfg.ExtensionsToIgnore) == 0 {
			cfg.ExtensionsToIgnore = ExtensionsToIgnore
		}
	}

	return func(c *fiber.Ctx) {
		if !shouldShowPrerenderedPage(c, cfg) {
			c.Next()
			return
		}

		c.Next()
	}
}

func shouldShowPrerenderedPage(c *fiber.Ctx, cfg PrerenderConfig) bool {
	baseURL := strings.ToLower(c.BaseURL())
	userAgent := c.Get("user-agent")

	if userAgent == "" {
		return false
	}
	if c.Method() != http.MethodGet && c.Method() != http.MethodHead {
		return false
	}
	if c.Get("x-prerender") != "" {
		return false
	}
	//if it contains _escaped_fragment_, show prerendered page
	if strings.Contains(c.OriginalURL(), "_escaped_fragment_") {
		return true
	}
	//if it is a bot...show prerendered page
	for _, bot := range cfg.CrawlerUserAgents {
		if strings.Contains(userAgent, bot) {
			return true
		}
	}
	//if it is BufferBot...show prerendered page
	if c.Get("x-bufferbot") != "" {
		return true
	}
	//if it is a bot and is requesting a resource...dont prerender
	for _, ext := range cfg.ExtensionsToIgnore {
		if strings.HasSuffix(baseURL, ext) {
			return false
		}
	}
	// TODO: check for whitelisted URL
	// TODO: check for blacklisted URL
	return true
}

func getPrerenderResponse(c *fiber.Ctx) error {
	// TODO
	return nil
}

func plainResponse(c *fiber.Ctx) error {
	// TODO
	return nil
}

func gzipResponse(c *fiber.Ctx) error {
	// TODO
	return nil
}
