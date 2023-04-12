package html

var htmlTmpl string = `
<html>

<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: "Helvetica", sans-serif;
            margin: 16px 64px
        }

        h4 {
            padding: 4px 8px;
        }


        .coverage {
            font-size: small;
            font-weight: normal;
        }

        a {
            padding: 8px;
            color: black;
            text-decoration: none;
            display: block;
            border: thin #eeeeee solid;
        }


        .passed a {
            background-color: #dfd;
        }

        .skipped a {
            background-color: #eee;
        }

        .failed a {
            background-color: #fdd;
        }

        .badge {
            background-color: #888;
            color: white;
            border-radius: 4px;
            padding: 2px 8px;
            float: right;
        }

        .passed .badge {
            background-color: #5c5;
        }

        .skipped .badge {}

        .failed .badge {
            background-color: #d66;
        }

        .expando {
            display: none;
        }

        .content {
            margin-left: 1em;
            border-bottom: thin #eeeeee solid;
            color: #444;
            white-space: pre-wrap;
        }

        .duration {
            font-size: x-small;
        }

        .expando:target {
            display: block
        }

        ul {
            list-style-type: none;
            margin: 0;
            padding: 0;
            overflow: hidden;
        }

        li {
            float: left;
        }

        li a {
            display: block;
            text-align: center;
            text-decoration: none;
        }

        .centered {
            justify-items: center;
            justify-content: center;
            align-items: stretch;
        }

        .pad {
            padding-bottom: 1em;
        }

        .card {
            display: flex;
            align-content: space-around;
            justify-content: space-around;
        }

        .container {
            border-right: 1px solid rgb(172, 172, 172);
            border-left: 1px solid rgb(172, 172, 172);
            margin-bottom: 1em;
        }
    </style>
</head>
{{range .}}
<div class="container">
    {{ $timestamp := .Timestamp }}
    {{ $path := .Path }}
    <div class="card">
        <h5>{{.Coverage}}% Coverage</h5>
        <h3>{{$path}}</h3>
        <h5>{{$timestamp |formatTime}}</h5>
    </div>
    <hr>
    {{range .Results}}
    <div
        class='{{if or .Error (not (eq .ResourceCountDiff 0))}}failed{{else if eq .ResourceCount 0}}skipped{{else}}passed{{end}} pad'>
        <a href='#{{.Path | strip}}'>
            {{.Path}}
            <span class='badge'>
                {{if .Error}}Error{{else if (not (eq .ResourceCountDiff 0))}}Fail{{else if eq
                .ResourceCount 0}}Skip{{else}}Pass
                {{end}}
            </span>
        </a>

        <ul class="centered">
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>Coverage:
                    {{.Coverage}}%</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ResourceCount:
                    {{.ResourceCount}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ResourceCountExists:
                    {{.ResourceCountExists}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ResourceCountDiff:
                    {{.ResourceCountDiff}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ActionNoopCount:
                    {{.ActionNoopCount}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ActionCreateCount:
                    {{.ActionCreateCount}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ActionReadCount:
                    {{.ActionReadCount}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ActionUpdateCount:
                    {{.ActionUpdateCount}}</a>
            </li>
            <li>
                <a class="flexed" href='#{{.Path | strip}}'>ActionDeleteCount:
                    {{.ActionDeleteCount}}</a>
            </li>
        </ul>


        <div class='expando' id='{{.Path | strip}}'>
            <div class='content'>
                <code>
                                {{.PlanRaw}}
                            </code>
            </div>
            <p class='duration' title='Test duration'>{{.Duration}}</p>
        </div>
    </div>
    {{end}}
</div>
{{end}}

</html>
</body>

</html>
`
