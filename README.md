# Serve Static

<!-- > 2018-07-04T16:59:21+0800 -->

<!-- Titles: *Serve Static*. -->

A library serving static files, like `http.FileServer` and `http.ServeFile`,
but in the *no-trailing-slash* version.

Currently static web servers are often and inherently involved with the trailing slash '/' in URL --
which has more to do with folders in static sites,
while the web servers hosting dynamic contents tend to get rid of the trailing slash
because a path without the trailing slash means a resource in nature following the REST rules.

This project is trying to host static sites without the trailing slash.

## How does it Work?

0. Serve the index.html for the home page '/'.
	- `/`
		- Serve `./index.html`
1. Redirect all trailing-slash paths to no-trailing-slash paths.
	- `/${folder-name}/${target-name}/`
		- Redirect to `/${folder-name}/${target-name}`
2. The \*.html files are not served.
3. Serve regular files(, but not \*.html).
	- `/${folder-name}/${target-name}/image-a.png`
		- Serve `./${folder-name}/${target-name}/image-a.png`
4. Serve the target html file if no dot is found in the request.path.
	- `/${folder-name}/${target-name}`
		- Serve `./${folder-name}/${target-name}/${target-name}.html` if any
		- Serve `./${folder-name}/${target-name}/index.html` if any
5. Pass through with the resolved path.

## Structure of Static Sites Hosted by the Library

We match the file `./site-root/some-section/some-content/some-content.html` for the path `/some-section/some-content` to get rid of the trailing slash,
so the structure of static sites hosted by this library are slightly different from normal sites --
the `.../a-folder/a-folder.html` will be matched to the path `.../folder`
and the `.../a-folder/index.html` will be not matched to the path `.../folder/` anymore.

The structure of contents of a static sites may like:

- `/index.html` The site home page.
- `/blogs` The blogs module.
	- `/any-folder`
		- `/a-normal-blog`
			- `/a-normal-blog.html` The blog, which will be matched with the path: `/blogs/any-folder/a-normal-blog`.
	- `blogs.html` The home page for the blogs module, which will be matched with the `/blogs` path.
- `/about` The about module.
	- `/about.html` The home page for the about module, which will be matched with the `/about` path.
		- The relative links like to `some-images.png` will break.
	- `/some-images.png` Images used in the about module.
	- `/some-attachments.zip` Attachments used in the about module.
- `/contact` The contact-us module.
	- `/contact.html` The home page for the contact-us module, which will be matched with the `/contact` path.

## Example of Usage

You may check out an application here: https://github.com/zhanbei/static-server --
A static server in Go, supporting hosting static files in the *no-trailing-slash* version.

Here is a basic example shown below, which can also be found here [./example/main.go](example/main.go).

```go
// @see ./example/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/zhanbei/serve-static"
)

var mStaticServer *servestatic.FileServer

const mAddress = "0.0.0.0:1234"
const mRootDir = "./static"
const mVirtualHost = false

func init() {
	server, err := servestatic.NewFileServer(mRootDir, mVirtualHost)
	if err != nil {
		panic(err)
	}
	mStaticServer = server
}

func main() {
	fmt.Println("Server is listening:", mAddress, "and looking after:", mRootDir, "; Using virtual host:", mVirtualHost, ".")
	// @see https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	http.ListenAndServe(mAddress, mStaticServer)
}
```

## `WARNING` The Usage of `<base>`, Relative Links and In-page Anchors

For a normal static sites, the relative links used will break since the trailing slash has been removed.

The usage of `<base href="${page-name}/">` will make the relative links work again while
[break the in-page anchors](https://stackoverflow.com/questions/8108836/make-anchor-links-refer-to-the-current-page-when-using-base)
however.

You have to decide whether to use the `<base href="${page-name}/">` or not yourself:

- Use a `<base>`
	- All in-page anchors will break.
- No `<base>`
	- All relative link will break.

*Recommend: No `<base>` and no relative links.*
