package main

import (
	"html/template"
	"log"
	"path"
)

func loadTemplates(dir string) *template.Template {
	if dir == "" {
		return getTemplates()
	}
	log.Printf("using custom template directory %q", dir)
	t, err := template.New("").ParseFiles(path.Join(dir, "sign_in.html"), path.Join(dir, "error.html"))
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}
	return t
}

func getTemplates() *template.Template {
	t, err := template.New("foo").Parse(`
{{define "sign_in.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">

<head>
  <title>Log In</title>
  <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;500;600;700&family=Raleway&display=swap"
    rel="stylesheet">
  <style type="text/css">
    :root {
      --xs: 320px;
      --sm: 481px;
      --md: 769px;
      --lg: 1366px;
      --xl: 1920px;
      --stakaterPrimary: #fd7401;
      --stakaterSecondary: #3e0054;
      --contentFont: 'Raleway', sans-serif;
      --headingFont: 'Montserrat', sans-serif;
      --space-1: .5rem;
      --space-2: 1rem;
      --space-3: 2rem;
      --space-4: 4rem;
      --error: #f44336;
      --warning: #ff9800;
      --success: #4caf50;
      --info: #2196f3;
    }

    html {
      -webkit-box-sizing: border-box;
      -moz-box-sizing: border-box;
      box-sizing: border-box;
      font-size: 14px;
    }

    *,
    *:before,
    *:after {
      -webkit-box-sizing: inherit;
      -moz-box-sizing: inherit;
      box-sizing: inherit;
    }

    label {
      display: block;
      margin-bottom: var(--space-1);
      font-weight: 400;
      text-transform: uppercase;
    }

    input {
      width: 100%;
      border-radius: 0;
      transition: .25s linear outline, .25s linear border;
      outline: 8px solid transparent;
      border: 1px solid var(--stakaterSecondary);
      padding: 10px 14px;
    }

    input:hover {
      border-color: var(--stakaterPrimary);
    }

    input:focus {
      outline: 2px solid var(--stakaterPrimary);
      border-color: var(--stakaterPrimary);
    }

    input:disabled {
      width: 100%;
      padding: .25rem .25rem;
      border-radius: 0;
      cursor: not-allowed;
      background: rgba(0, 0, 0, .01);
      border: 1px solid rgba(0, 0, 0, .25);
    }


    .flat-button {
      position: relative;
      border: none;
      min-width: 100px;
      font-size: 12pt;
      outline: none;
      padding: 6px 18px;
      box-shadow: none;
      text-transform: uppercase;
    }

    .flat-button:focus:before,
    .flat-button:hover:before,
    .flat-button:active:before {
      display: inline-block;
      content: "";
      width: 100%;
      height: 100%;
      background-color: rgba(0, 0, 0, 0.1);
      position: absolute;
      left: 0;
      top: 0;
    }

    .flat-button:active:focus:before {
      background-color: rgba(0, 0, 0, 0.2);
    }

    .flat-button-primary {
      background: var(--stakaterPrimary);
      color: white;
    }

    .flat-button-lg {
      padding: 8px 18px;
      width: 100%;
    }

    body {
      font-family: Montserrat, Helvetica, Arial, sans-serif;
      font-weight: 400;
      line-height: 1.43;
      background: linear-gradient(220.21deg, #3e0054 15%, #fd7401 75.99%, #f89a02);
      min-height: 100vh;
      font-size: 14px;
      color: #444444;
    }

    .overlay {
      display: flex;
      flex-direction: column;
      align-items: center;
      height: 100vh;
      justify-content: center;
    }

    .overlay::before {
      position: absolute;
      display: block;
      content: "";
      width: 100%;
      height: 100%;
      top: 0px;
      left: 0px;
      background: url("https://cloud.stakater.com/b529c3b0853a01afb49dad546a652735.png") 0px 0px / cover no-repeat padding-box padding-box transparent;
      opacity: 0.12;
      pointer-events: none;
    }

    .form-container {
      width: 80%;
      max-width: 500px;
      background-color: white;
      display: flex;
      align-items: space-evenly;
    }

    .form-block-container {
      display: flex;
      align-items: space-evenly;
    }

    .form-block {
      padding: var(--space-4);
      width: 100%;
    }

    .error-container {
      padding: var(--space-4);
      color: var(--error);
    }

    .logo {
      display: flex;
      font-size: 24px;
      white-space: nowrap;
      margin-bottom: 50px;
      justify-content: center;
      color: black;
    }

    .logo svg {
      height: 58px;
    }

    .logo .product {
      text-align: left;
    }

    .submit-block {
      align-items: center;
      color: white;
      text-align: center;
    }
  </style>
</head>
<body>
  <div class="overlay">
    <div class="form-container">
      <div class="form-block submit-block">
        <div class="logo">
          <svg xmlns="http://www.w3.org/2000/svg" with="56" height="56" viewBox="0 0 78 59.841"
            preserveAspectRatio="xMidYMid meet">
            <g id="Group_22" transform="translate(-16.176 -14.28)">
              <path id="Path_26"
                d="M20.931,16.88l20.486-.071L58.339,50.544l.006,1.442-17.22.06L41.1,45.951l-2.922-2.167-17.161.06Zm-2.248-2.465.145,41.838,42.033-.145-.021-6.514L43.293,14.331Z"
                transform="translate(1.527 0.031)" />
              <rect id="Rectangle_37" width="42.034" height="11.147" transform="translate(20.354 56.283) rotate(-0.197)"
                fill="#fecd07" />
              <path id="Path_27" d="M45.687,14.3l.182,52.983L60,67.23l14.114-4.052-24.254.083-.168-48.98Z"
                transform="translate(17.97 0.001)" />
              <path id="Path_28" d="M24.995,56.91a7.712,7.712,0,1,1,7.686-7.739,7.721,7.721,0,0,1-7.686,7.739"
                transform="translate(0.658 16.565)" />
              <path id="Path_29"
                d="M25.19,42.378a7.064,7.064,0,1,1-7.041,7.089,7.065,7.065,0,0,1,7.041-7.089m0-1.3a8.359,8.359,0,1,0,8.389,8.331,8.367,8.367,0,0,0-8.389-8.331"
                transform="translate(0.413 16.32)" fill="#fff" />
              <path id="Path_30" d="M41.423,56.852a7.712,7.712,0,1,1,7.686-7.737,7.718,7.718,0,0,1-7.686,7.737"
                transform="translate(10.662 16.531)" />
              <path id="Path_31"
                d="M41.618,42.321a7.065,7.065,0,1,1-7.039,7.089,7.067,7.067,0,0,1,7.039-7.089m0-1.3A8.36,8.36,0,1,0,50,49.357a8.372,8.372,0,0,0-8.389-8.331"
                transform="translate(10.416 16.286)" fill="#fff" />
              <path id="Path_32" d="M39.578,50.19A2.878,2.878,0,1,1,42.447,47.3a2.879,2.879,0,0,1-2.869,2.888"
                transform="translate(12.491 18.36)" fill="#fff" />
              <path id="Path_33"
                d="M39.651,44.767a2.636,2.636,0,1,1-2.627,2.645,2.635,2.635,0,0,1,2.627-2.645m0-.484a3.12,3.12,0,1,0,3.131,3.11,3.123,3.123,0,0,0-3.131-3.11"
                transform="translate(12.399 18.269)" fill="#fff" />
              <path id="Path_34" d="M23.185,50.246a2.878,2.878,0,1,1,2.869-2.888,2.88,2.88,0,0,1-2.869,2.888"
                transform="translate(2.509 18.394)" fill="#fff" />
              <path id="Path_35"
                d="M23.259,44.823a2.636,2.636,0,1,1-2.629,2.645,2.638,2.638,0,0,1,2.629-2.645m0-.484a3.12,3.12,0,1,0,3.131,3.11,3.123,3.123,0,0,0-3.131-3.11"
                transform="translate(2.418 18.303)" fill="#fff" />
              <path id="Path_36" d="M16.176,38.412a3.964,3.964,0,1,1,3.977,3.951,3.963,3.963,0,0,1-3.977-3.951"
                transform="translate(0 12.272)" />
              <path id="Path_37" d="M23.585,36.866,22.008,23.448l1.826-.006L26.1,33.09l8.008.5.011,3.237Z"
                transform="translate(3.551 5.578)" fill="#fecd07" />
              <path id="Path_38"
                d="M22.032,41.083c5.327-.019,10.221,2.232,11.219,9.969l-2.989.011c.806-6.576-3.461-9.652-8.23-9.98"
                transform="translate(3.566 16.32)" fill="#deb306" />
              <path id="Path_39"
                d="M42.681,41.026c-5.329.019-10.205,2.3-11.151,10.048l2.989-.011c-.851-6.568,3.4-9.674,8.162-10.036"
                transform="translate(9.349 16.286)" fill="#deb306" />
              <path id="Path_40"
                d="M18.756,35.529,36.7,35.468l1.093.753.024,6.857,20.486-2.239L41.087,40.9,41.066,34.8l-2.922-2.168-19.4.067Z"
                transform="translate(1.565 11.178)" fill="#333" />
              <path id="Path_41" d="M40.256,42.667A6.794,6.794,0,0,1,43,51.336a9.274,9.274,0,0,0-2.746-8.669"
                transform="translate(14.663 17.285)" fill="#fff" />
              <path id="Path_42" d="M23.713,42.724a6.789,6.789,0,0,1,2.746,8.667,9.271,9.271,0,0,0-2.746-8.667"
                transform="translate(4.589 17.32)" fill="#fff" />
              <path id="Path_43" d="M23.323,34.9l10.016-.034.817.581.022,6.272-.935-5.586-.391-.467Z"
                transform="translate(4.352 12.536)" fill="#fff" />
              <path id="Path_44" d="M22.751,23.444,24.66,35.079l9-.654,0-.833-8.008-.5-2.27-9.648Z"
                transform="translate(4.004 5.579)" fill="#bf9a05" />
              <path id="Path_45" d="M51.312,36.1l-.133-.126,0,0Z" transform="translate(21.313 13.207)" fill="#957552" />
              <path id="Path_46" d="M51.007,24.922l-.116-.108,0,0Z" transform="translate(21.139 6.412)"
                fill="#957552" />
              <rect id="Rectangle_38" width="18.211" height="18.211" transform="translate(69.572 43.865) rotate(-0.197)"
                fill="#c69c6d" />
              <rect id="Rectangle_39" width="3.912" height="3.955" transform="translate(76.721 43.84) rotate(-0.197)"
                fill="#957552" />
              <rect id="Rectangle_40" width="3.912" height="3.956" transform="translate(76.77 58.095) rotate(-0.197)"
                fill="#957552" />
              <rect id="Rectangle_41" width="15.708" height="15.708" transform="translate(69.513 26.641) rotate(-0.197)"
                fill="#c69c6d" />
              <rect id="Rectangle_42" width="3.375" height="3.413" transform="translate(75.678 26.619) rotate(-0.197)"
                fill="#957552" />
              <rect id="Rectangle_43" width="3.375" height="3.413" transform="translate(75.721 38.914) rotate(-0.197)"
                fill="#957552" />
              <rect id="Rectangle_44" width="10.287" height="10.286" transform="translate(69.47 14.316) rotate(-0.197)"
                fill="#c69c6d" />
              <rect id="Rectangle_45" width="2.21" height="2.234" transform="translate(73.526 14.565)" fill="#957552" />
              <rect id="Rectangle_46" width="2.21" height="2.234" transform="translate(73.537 22.354) rotate(-0.197)"
                fill="#957552" />
            </g>
          </svg>
          <span class="product">Stakater <br> App Agility Platform</span>
        </div>

    <form method="GET" action="{{.ProxyPrefix}}/start">
        <input type="hidden" name="rd" value="{{.Redirect}}">
        {{ if .SignInMessage }}
          <p>{{.SignInMessage}}</p>
        {{ end}}
        <button type="submit" class="flat-button flat-button-primary">Log in</button>
      </form>
    </div>
  </div>
    </div>
  <script>
    if (window.location.hash) {
      (function() {
        var inputs = document.getElementsByName('rd');
        for (var i = 0; i < inputs.length; i++) {
          inputs[i].value += window.location.hash;
        }
      })();
    }
  </script>
</body>
</html>
{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}

	t, err = t.Parse(`{{define "error.html"}}
  <!DOCTYPE html>
  <html>
  <head>
    <title>Error</title>
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;500;600;700&family=Raleway&display=swap"
      rel="stylesheet">
    <style type="text/css">
      :root {
        --xs: 320px;
        --sm: 481px;
        --md: 769px;
        --lg: 1366px;
        --xl: 1920px;
        --stakaterPrimary: #fd7401;
        --stakaterSecondary: #3e0054;
        --contentFont: 'Raleway', sans-serif;
        --headingFont: 'Montserrat', sans-serif;
        --space-1: .5rem;
        --space-2: 1rem;
        --space-3: 2rem;
        --space-4: 4rem;
        --error: #f44336;
        --warning: #ff9800;
        --success: #4caf50;
        --info: #2196f3;
      }
  
      html {
        -webkit-box-sizing: border-box;
        -moz-box-sizing: border-box;
        box-sizing: border-box;
        font-size: 14px;
      }
  
      *,
      *:before,
      *:after {
        -webkit-box-sizing: inherit;
        -moz-box-sizing: inherit;
        box-sizing: inherit;
      }
  
      body {
        font-family: Montserrat, Helvetica, Arial, sans-serif;
        font-weight: 400;
        line-height: 1.43;
        background: linear-gradient(220.21deg, #3e0054 15%, #fd7401 75.99%, #f89a02);
        min-height: 100vh;
        font-size: 14px;
        color: #444444;
      }
  
      .overlay {
        display: flex;
        flex-direction: column;
        align-items: center;
        height: 100vh;
        justify-content: center;
      }
  
      .overlay::before {
        position: absolute;
        display: block;
        content: "";
        width: 100%;
        height: 100%;
        top: 0px;
        left: 0px;
        background: url("https://cloud.stakater.com/b529c3b0853a01afb49dad546a652735.png") 0px 0px / cover no-repeat padding-box padding-box transparent;
        opacity: 0.12;
        pointer-events: none;
      }
  
      .content {
        display: flex;
        flex-direction: column;
        background-color: white;
        padding: 3rem;
        width: 80%;
        max-width: 700px;
        margin: auto;
      }
    </style>
  </head>
  
  <body>
    <div class="overlay">
      <div class="content">
          <h1>{{.Title}}</h1>
          <p>{{.Message}}</p>
          <p><a href="{{.ProxyPrefix}}/sign_in">Log In</a></p>    </div>
    </div>
  
  </body>
  
  </html>
{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}
	return t
}
