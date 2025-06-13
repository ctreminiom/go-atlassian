package confluence

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ChildrenDescendantConnector interface {

	// Children returns a map of the direct children of a piece of content.
	//
	// A piece of content has different types of child content, depending on its type.
	//
	// These are the default parent-child content type relationships:
	//
	// page: child content is page, comment, attachment
	//
	// blogpost: child content is comment, attachment
	//
	// attachment: child content is comment
	//
	// comment: child content is attachment
	//
	// GET /wiki/rest/api/content/{contentID}/child
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-children
	Children(ctx context.Context, contentID string, expand []string, parentVersion int) (*model.ContentChildrenScheme, *model.ResponseScheme, error)

	// Move moves a page from its current location in the hierarchy to another.
	//
	// Position describes where in the hierarchy the page should be moved to in
	// relationship to targetID.
	//
	// before: page will be a sibling of target but show up just before target in
	// the list of children
	//
	// after: page will be a sibling of target but show up just after target in the
	// list of children
	//
	// append: page will be a child of the target and be appended to targets list of
	// children
	//
	// PUT /wiki/rest/api/content/{contentID}/move/{position}/{targetID}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#move
	Move(ctx context.Context, contentID string, position string, targetID string) (*model.ContentMoveScheme, *model.ResponseScheme, error)

	// ChildrenByType returns all children of a given type, for a piece of content.
	//
	// A piece of content has different types of child content
	//
	// GET /wiki/rest/api/content/{contentID}/child/{type}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-children-by-type
	ChildrenByType(ctx context.Context, contentID, contentType string, parentVersion int, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error)

	// Descendants returns a map of the descendants of a piece of content.
	//
	// This is similar to Get content children, except that this method returns child pages at all levels,
	//
	// rather than just the direct child pages.
	//
	// GET /wiki/rest/api/content/{contentID}/descendant
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-descendants
	Descendants(ctx context.Context, contentID string, expand []string) (*model.ContentChildrenScheme, *model.ResponseScheme, error)

	// DescendantsByType returns all descendants of a given type, for a piece of content.
	//
	// This is similar to Get content children by type,
	//
	// except that this method returns child pages at all levels, rather than just the direct child pages.
	//
	// GET /wiki/rest/api/content/{contentID}/descendant/{type}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-descendants-by-type
	DescendantsByType(ctx context.Context, contentID, contentType, depth string, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error)

	// CopyHierarchy copy page hierarchy allows the copying of an entire hierarchy of pages and their associated properties,
	//
	// permissions and attachments. The id path parameter refers to the content id of the page to copy,
	//
	// and the new parent of this copied page is defined using the destinationPage id in the request body.
	//
	// The titleOptions object defines the rules of renaming page titles during the copy;
	//
	// for example, search and replace can be used in conjunction to rewrite the copied page titles.
	//
	// RESPONSE =  Use the /longtask/ REST API to get the copy task status.
	//
	// POST /wiki/rest/api/content/{contentID}/pagehierarchy/copy
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#copy-page-hierarchy
	CopyHierarchy(ctx context.Context, contentID string, options *model.CopyOptionsScheme) (*model.TaskScheme, *model.ResponseScheme, error)

	// CopyPage copies a single page and its associated properties, permissions, attachments, and custom contents.
	//
	// The id path parameter refers to the content ID of the page to copy.
	//
	// The target of the page to be copied is defined using the destination in the request body and can be one of the following types.
	//
	// 1. space: page will be copied to the specified space as a root page on the space
	//
	// 2. parent_page: page will be copied as a child of the specified parent page
	//
	// 3. existing_page: page will be copied and replace the specified page
	//
	// By default, the following objects are expanded: space, history, version.
	//
	// POST /wiki/rest/api/content/{contentID}/copy
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#copy-single-page
	CopyPage(ctx context.Context, contentID string, expand []string, options *model.CopyOptionsScheme) (*model.ContentScheme, *model.ResponseScheme, error)
}
