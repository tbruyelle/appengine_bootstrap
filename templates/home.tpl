{{define "content"}}
<div class="container">
  <div class="row">
    {{ if .User }}
    <table class="table table-striped">
        <thead>
            <tr><th>Name</th><th>date</th></tr>
        </thead>
        <tbody>
        {{range .Registrations}}
            <tr><td>{{.Name}}</td><td>{{.Date}}</td></tr>
        {{end}}
        </tbody>
    </table>
    {{else}}
    <div class="jumbotron">
      <h1>Hello, World!</h1>
      <p>Welcome to your new Go web project.</p>
    </div>
    {{end}}
  </div>
</div>
{{end}}
