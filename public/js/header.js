$(document).ready(function() {
  authed = document.getElementById("authed-bar");
  notAuthed = document.getElementById("not-authed-bar");

  if (isUserAuthed()) {
    authed.classList.remove("hidden");
  } else {
    notAuthed.classList.remove("hidden");
  }
});
