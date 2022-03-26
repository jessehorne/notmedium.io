$(document).ready(function() {
  if (!isUserAuthed()) {
    window.location.href = "/";
  }

  // fill data
  var user = getUser();

  $("#username").html(user.Username);
  $("#createdat").html(user.CreatedAt);

  // handle theme stuff
  $("#theme").html(window.localStorage.getItem("theme"));

  $("#setlight").click(function() {
    window.localStorage.setItem("theme", "light");
    window.location.href = "/settings";
  });

  $("#setdark").click(function() {
    window.localStorage.setItem("theme", "dark");
    window.location.href = "/settings";
  });

  console.log("settings...");
});
