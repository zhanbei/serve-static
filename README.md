# Serve Static

<!-- > 2018-07-04T16:59:21+0800 -->

<!-- Titles: *Serve Static*. -->

A library serving static files, like `http.FileServer` and `http.ServeFile`, but in the no-trailing-slash version.

Currently static web servers are often and inherently involved with the trailing slash in URL -- the trailing slash has more to do with folders in static sites, while the web servers hosting dynamic contents tend to get rid of the trailing slash because a path without the trailing slash means a resource.

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
		- Serve `./${folder-name}/${target-name}/${target-name}.html`
5. Pass through with the resolved path.
