function newLightTheme() {
  return {
    "bg": "#fff",
    "text": "#000",
    "links": "#00a",
    "inputBG": "#fff",
    "inputText": "#000",
    "buttonBG": "#fff",
    "buttonText": "#000",
  }
}

function newDarkTheme() {
  return {
    "bg": "#000",
    "text": "#fff",
    "links": "#00a",
    "inputBG": "#000",
    "inputText": "#fff",
    "buttonBG": "#007",
    "buttonText": "#fff",
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
  $("input").css("background-color", t.inputBG);
  $("input").css("color", t.inputText);
  $("textarea").css("background-color", t.inputBG);
  $("textarea").css("color", t.inputText);
  $("button").css("background-color", t.buttonBG);
  $("button").css("color", t.buttonText);
});
