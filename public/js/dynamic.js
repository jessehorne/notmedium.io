function newLightTheme() {
  return {
    "bg": "#fff",
    "text": "#000",
    "links": "#00a"
  }
}

function newDarkTheme() {
  return {
    "bg": "#000",
    "text": "#fff",
    "links": "#00a"
  }
}


$(document).ready(function() {
  // check if theme settings exist
  var theme = window.localStorage.getItem("theme");

  if (theme == null) {
    window.localStorage.setItem("theme", "light");
    theme = "light";
  }

  var t;

  if (theme == "light") {
    t = newLightTheme();
  } else if (theme == "dark") {
    t = newDarkTheme();
  }

  $("body").css("background-color", t.bg);
  $("body").css("color", t.text);
  $("a").css("color", t.links);
});
