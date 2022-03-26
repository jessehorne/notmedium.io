$(document).ready(function() {
  if (!isUserAuthed()) window.location.href = "/";

  var userID = $("#userID").val();

  // fetch article
  (async () => {
    const raw = await fetch('http://localhost:8080/api/users/' + userID, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getUser().ApiToken,
      }
    });

    const res = await raw.json();

    if (raw.status === 200) {
      var user = res.data.user;
      var articles = res.data.articles

      $("#username").html(user.Username);
      $("#createdat").html(user.CreatedAt);

      // load articles
      if (articles.length > 0) {
        articles.forEach(function(a) {
          var newString = '<div><a href="/a/ID">TITLE</a> on CREATEDAT</div>';

          newString = newString.replace("TITLE", a.Title);

          var d = new Date(a.CreatedAt);
          var datestring = (d.getMonth()+1) + "-" + d.getDate() + "-" + d.getFullYear();

          newString = newString.replace("CREATEDAT", datestring);

          var newElem = $(newString);

          $("#articles").append(newElem);
        });
      } else {
        $("#articles").html("There aren't any articles published at the moment.");
      }
    } else {
      console.log("not 200")
    }

    console.log("done");

  })();
});
