<p align="center">
  <a href="https://ghatd.com">
    <img alt="GHATD" src="https://demo.ghatd.com/static/img/ghat-logo-square.png" width="180" />
  </a>
</p>
<h1 align="center">
  GHAT(D)'s <code>Web Detail</code> - Vite Vue (TS) Template [WIP]
</h1>
<h4 align="center">
  Apart of the <a href="https://github.com/ooaklee/ghatd/pull/2" target="_blank">GHAT(D) initiative</a>
</h4>


> This repo is strongly linked to this [**GHAT(D) PR** being merged](https://github.com/ooaklee/ghatd/pull/2)

Use GHAT(D) `Details` to hit the floor running with your next Go-backed full-stack web application. This `Detail` has everything needed to get started with a Vue frontend (based on [**huibizhang's vite vue** template](https://github.com/huibizhang/template-vite-vue-ts-tailwind-v3)) which will be embedded in your final go binary. 

Use this to Kick off your project if you want your application's **web `Detail`** to be a Vue-based frontend.


## 🚥 Getting started

`Details` are independent applications that can run within the GHAT(D) framework and individually. To run this `web` Detail locally, please:

- Ensure you have the appropriate version of Go installed
- Ensure you have the correct versions of  **node** and **vue** installed
  - This can be done using `asdf install`
- Install the node packages using `yarn install`
- There are multiple ways to run your code when testing locally.
  - To run your code with host reloading, use `yarn dev`
  - To see how your code will run in the embedded binary, run `yarn build && yarn go:web:start`
  - To run your code from the compiled binary, run: `yarn go:web:build-arm64`

- If you built your frontend code, `yarn build` you should be able to access the `web` Detail on http://localhost:4044. Otherwise, it's recommended to use `yarn dev` which will require hot reloading from Vitril.


For the best experience developing this Detail, we recommend using hot reloading with Vite by using the following:
```sh
yarn dev
```

## 🪡 Putting together your web application and the ghat(d) framework (TBC)

1.  **Create a GHAT(D) Web App.**

 Use the GHATDCLI to create a new web app, specifying this `web'detail.

 ```shell
    # create a new GHAT(D) Web App using this demo web detail
    ghatdcli new -n "my-new-web-app" -m "github.com/some-org-or-personal/my-new-web-app" -w "https://github.com/ooaklee/ghatd-detail-web-vite-vue-ts-template"
 ```

2.  **Start developing.**

Navigate into your new web app's directory and start it up.

```shell
cd my-new-web-app/
asdf install
go mod tidy

yarn start # runs "go run main.go start-server" and "yarn dev" concurrently

## OR ##

yarn start-hot # runs "reflex -r '\.(html|go|css|png|svg|ico|js|woff2|woff|ttf|eot)$' -s -- go run main.go start-server" and "yarn dev" concurrently
```

> For the best experience, we recommend using hot reloading with the server when developing by using `reflex -r '\.(html|go|css|png|svg|ico|js|woff2|woff|ttf|eot)$' -s -- go run main.go start-server`. You will have to ensure you have 
[**reflex** (click to go to installation steps)](https://github.com/cespare/reflex?tab=readme-ov-file#installation) installed.


3.  **Browse the service and start editing your app code!**

The backend server will now be running at `http://localhost:4000`, but with this `web` detail, the frontend will be accessible on `http://localhost:5173/`.

Note: The frontend has a proxy set up, `http://localhost:5173/api`, that points to the back end. This is done so that when integrating with the backend, you can use the URI `/api/...`, which will map to the backend, and the code will work when everything is compiled and embedded into a single binary


## More about this `web` Detail

See below for more information on the core components used for this `Detail`'s stack.

- **go:** [v1.22.x](https://go.dev/doc/install)
- **vue:** [^3.2.25](https://vuejs.org/)
- **vite:** [v2.x](https://v2.vitejs.dev/)
- **tailwindcss:** [v3.x](https://github.com/asdf-community/asdf-golang)
- **daisy ui:** [^4.12.13](https://daisyui.com/docs/install/)
  - Notable alternatives include:
    - **flowbite:** [v2.3.x](https://flowbite.com/docs/getting-started/introduction/#include-via-cdn)
    - **wind-ui:** [v.3.4.x](https://wind-ui.com/)
- **version manager:** [asdf](https://github.com/asdf-vm/asdf)



## License
This project is licensed under the [MIT License](./LICENSE).
