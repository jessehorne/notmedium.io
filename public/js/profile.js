function deleteArticle(id) {
  (async () => {
    const raw = await fetch('http://localhost:8080/api/articles/' + id, {
      method: 'DELETE',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getUser().ApiToken
      }
    });

    if (raw.status === 200) {
      window.location.href = "/profile";
    } else {
      window.location.href = "/profile";
    }

  })();
}

function loadArticles() {
  $("#articles").html("");

  var userID = getUser().ID;

  (async () => {
    const raw = await fetch('http://localhost:8080/api/users/' + userID + '/articles', {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getUser().ApiToken
      }
    });

    const res = await raw.json();

    var articles = res.data.articles;

    if (raw.status === 200) {
      if (articles.length > 0) {
        articles.forEach(function(a) {
          var newString = '<div><b>STATUS</b> | <a href="/a/ID">TITLE</a> on CREATEDAT | <a href="/a/ID/edit">Edit</a> <a href="#" class="delete" data-id="ID">Delete</a></div>';

          newString = newString.replace("STATUS", a.Published ? "Published" : "Draft")
          newString = newString.replace("TITLE", a.Title);
          newString = newString.replace("AUTHOR", a.Author);
          newString = newString.replace(/ID/g, a.ID);

          var d = new Date(a.CreatedAt);
          var datestring = (d.getMonth()+1) + "-" + d.getDate() + "-" + d.getFullYear();

          newString = newString.replace("CREATEDAT", datestring);

          var newElem = $(newString);

          $("#articles").append(newElem);
        });

        // delete article functionality
        $(".delete").click(function(e) {
          var articleID = $(this).data("id");

          deleteArticle(articleID);
        });
      } else {
        $("#articles").html("You haven't written anything yet!");
      }
    } else {
      $("#articles").html("You haven't written anything yet!");
    }

  })();
}

$(document).ready(function() {
  if (!isUserAuthed()) {
    window.location.href = "/";
  }

  // fill data
  var user = getUser();

  $("#username").html(user.Username);

  var then = new Date(user.CreatedAt);
  var formatted = ((then.getMonth() > 8) ? (then.getMonth() + 1) : ('0' + (then.getMonth() + 1))) + '/' + ((then.getDate() > 9) ? then.getDate() : ('0' + then.getDate())) + '/' + then.getFullYear()

  $("#createdat").html(formatted);

  // handle theme stuff
  $("#theme").html(window.localStorage.getItem("theme"));

  $("#setlight").click(function() {
    window.localStorage.setItem("theme", "light");
    window.location.href = "/profile";
  });

  $("#setdark").click(function() {
    window.localStorage.setItem("theme", "dark");
    window.location.href = "/profile";
  });

  // articles
  loadArticles();
});
