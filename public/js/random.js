function getFormData($form){
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function(n, i){
        indexed_array[n['name']] = n['value'];
    });

    return JSON.stringify(indexed_array);
}

function isUserAuthed() {
  // is user authed?
  var user = window.localStorage.getItem("user");

  if (user === null) {
    return false;
  }

  user = JSON.parse(user);

  var expired = (new Date(user.ApiTokenExpiresAt)) < (new Date());

  var loggedIn = !expired;

  return loggedIn;
}

function getUser() {
  return JSON.parse(window.localStorage.getItem("user"));
}

function logout() {
  window.localStorage.removeItem("user");

  window.location.href = "/";
}
