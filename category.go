package rsscast

// Category represents a podcast category with optional subcategories.
//
// Select the category that best reflects the content of your show.
// If available, you can also define a subcategory.
// Although you can specify more than one category and subcategory in your RSS feed,
// Apple Podcasts only recognizes the first category and subcategory.
//
// When specifying categories and subcategories, be sure to properly escape ampersands. For example:
//
// Single category:
//
//	<itunes:category text="History" />
//
// Category with ampersand:
//
//	<itunes:category text="Kids &amp; Family" />
//
// Category with subcategory:
//
//	<itunes:category text="Society &amp; Culture">
//	    <itunes:category text="Documentary" />
//	</itunes:category>
//
// Multiple categories:
//
//	<itunes:category text="Society &amp; Culture">
//	    <itunes:category text="Documentary" />
//	</itunes:category>
//	<itunes:category text="Health">
//	    <itunes:category text="Mental Health" />
//	</itunes:category>
//
// See Apple Podcast categories:
// https://podcasters.apple.com/support/1691-apple-podcasts-categories
type Category struct {
	Text          string
	Subcategories []string
}

// NewCategory creates a new Category with optional subcategories.
func NewCategory(text string, subcategories ...string) Category {
	return Category{
		Text:          text,
		Subcategories: subcategories,
	}
}
