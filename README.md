<h1>errorHandler_golang - A simple error handler for go</h1>
<h4> VERY EXPLICIT Error handler with mail by mailgun support for golang </h4>

<h2> Purpose </h2>
<p> The purpose of this error handler is, despite it's name, not to handle errors, but rather, notify admins and log, when errors occur. <br />
The logger is VERY explicit. </p>

<h3> What is errorHandler_golang not? </h3>
<p> 
    errorHandler_golang is NOT focused around speed, and is NOT going to handle hard use cases where security of the mailgun is an issue. <br />
    If you have staff on your team that should not have access to your mailgun key, etc. this software is not for you.
</p>