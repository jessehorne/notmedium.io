function loadArticles() {
  $("#articles").html("");

  (async () => {
    const raw = await fetch('http://localhost:8080/api/articles', {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getUser().ApiToken
      }
    });

    const res = await raw.json();

    var articles = res.data.articles;
    console.log(articles);
    // var count = res.data.count;
    // var limit = res.data.limit;
    // var page = res.data.page;

    if (raw.status === 200) {
      if (articles.length > 0) {
        articles.forEach(function(a) {
          var newString = '<div><a href="/a/ID">TITLE by AUTHOR</a></div>';

          newString = newString.replace("TITLE", a.Title);
          newString = newString.replace("AUTHOR", a.UserID);
          newString = newString.replace("ID", a.UserID);
          newString = newString.replace("CREATEDAT", a.CreatedAt);

          var newElem = $(newString);

          $("#articles").append(newElem);
        });
      } else {
        $("#articles").html("There aren't any articles published at the moment.");
      }
    } else {
      $("#articles").html("There aren't any articles published at the moment.");
    }

  })();
}

$(document).ready(function() {
  if (isUserAuthed()) {
    loadArticles();
  }
});
