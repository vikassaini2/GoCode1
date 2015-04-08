
define(["text!../../templates/loginPartial.htm"], function (loginPartial) {
    /*globals _:false, Handlebars:false, Backbone:false, esri: false, dojo: false */
    debugger;
    var compiled_template = Handlebars.compile(loginPartial);
    var data = {PageTitle: "Login Page"};
    $("#loginContainer").append(compiled_template(data));
});



