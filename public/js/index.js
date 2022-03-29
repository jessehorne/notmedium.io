const urlParams = new URLSearchParams(window.location.search);

var page = urlParams.get("page");
var limit = urlParams.get("limit");

function loadArticles(page) {

  (async (page) => {
    const raw = await fetch('http://localhost:8080/api/articles?page=' + page, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getUser().ApiToken
      }
    });

    const res = await raw.json();

    var articles = res.data.articles;

    // var count = res.data.count;
    // var limit = res.data.limit;
    // var page = res.data.page;

    if (raw.status === 200) {
      if (articles.length > 0) {
        articles.forEach(function(a) {
          var newString = '<div><a href="/a/ID">TITLE</a> by <a href="/profile/' + a.UserID + '">AUTHOR</a> on CREATEDAT</div>';

          newString = newString.replace("TITLE", a.Title);
          newString = newString.replace("AUTHOR", a.Author);
          newString = newString.replace("ID", a.ID);

          var d = new Date(a.CreatedAt);
          var datestring = (d.getMonth()+1) + "-" + d.getDate() + "-" + d.getFullYear();

          newString = newString.replace("CREATEDAT", datestring);

          var newElem = $(newString);

          $("#articles").append(newElem);
        });
      } else {
        // $("#articles").html("There aren't any articles published at the moment.");
      }
    } else if (raw.status == 401) {
      window.location.href = "/login";
    } else {
      // $("#articles").html("There aren't any articles published at the moment.");
    }

  })(page);
}

$(document).ready(function() {
  if (!isUserAuthed()) window.location.href = "/login";

  loadArticles(0);

  $("#load-more-link").click(function(e) {
    e.preventDefault();

    page = page + 1;
    loadArticles(page);
  });
});
