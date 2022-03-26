$(document).ready(function() {
  if (!isUserAuthed()) window.location.href = "/";

  var articleID = $("#articleID").val();

  // fetch article
  (async () => {
    const raw = await fetch('http://localhost:8080/api/articles/' + articleID, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getUser().ApiToken
      }
    });

    const res = await raw.json();

    if (raw.status === 200) {
      var article = res.data;

      $("#title").html(article.Title);
      $("#content").html(article.Content);

      console.log(article)

      if (article.Published) {
        // nothing
      } else {
        $("#ifDraft").removeClass("hidden");
      }
    } else {
      window.location.href = "/";
    }

    console.log("done");

  })();

});
