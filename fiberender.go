package fiberender

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber"
)

// PrerenderConfig can be used to customize Prerender
type PrerenderConfig struct {
	Skip               func(*fiber.Ctx) bool
	Token              string
	ServiceURL         string
	Host               string
	ForwardHeaders     bool
	Protocol           string
	CrawlerUserAgents  []string
	ExtensionsToIgnore []string
	Whitelist          []regexp.Regexp
	Blacklist          []regexp.Regexp

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
		if (cfg.Skip != nil && cfg.Skip(c)) || !shouldShowPrerenderedPage(c, cfg) {
			c.Next()
			return
		}

		c.Next()
	}
}

func shouldShowPrerenderedPage(c *fiber.Ctx, cfg PrerenderConfig) bool {
	baseURL := strings.ToLower(c.BaseURL())
	userAgent := c.Get("user-agent")
	referer := c.Get("referer")

	if userAgent == "" {
		return false
	}
	if c.Method() != http.MethodGet && c.Method() != http.MethodHead {
		return false
	}
	if c.Get("x-prerender") != "" {
		return false
	}

	shouldPrerender := false
	// if it contains _escaped_fragment_, show prerendered page
	if strings.Contains(c.OriginalURL(), "_escaped_fragment_") {
		shouldPrerender = true
	}
	// if it is a bot...show prerendered page
	for _, bot := range cfg.CrawlerUserAgents {
		if strings.Contains(userAgent, bot) {
			shouldPrerender = true
			break
		}
	}
	// if it is BufferBot...show prerendered page
	if c.Get("x-bufferbot") != "" {
		shouldPrerender = true
	}
	// if it is a bot and is requesting a resource...dont prerender
	for _, ext := range cfg.ExtensionsToIgnore {
		if strings.HasSuffix(baseURL, ext) {
			return false
		}
	}
	// if it is a bot and not requesting a resource and is not whitelisted...dont prerender
	if len(cfg.Whitelist) > 0 {
		inWhitelist := false
		for _, w := range cfg.Whitelist {
			if w.MatchString(baseURL) {
				inWhitelist = true
				break
			}
		}
		if !inWhitelist {
			return false
		}
	}
	// if it is a bot and not requesting a resource and is not blacklisted(url or referer)...dont prerender
	if len(cfg.Blacklist) > 0 {
		inBlacklist := false
		for _, w := range cfg.Blacklist {
			blacklistedURL := w.MatchString(baseURL)
			blacklistedReferer := w.MatchString(referer)
			if blacklistedURL || blacklistedReferer {
				inBlacklist = true
				break
			}
		}
		if inBlacklist {
			return false
		}
	}

	return shouldPrerender
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
