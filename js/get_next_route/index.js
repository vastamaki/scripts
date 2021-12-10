const fetch = require("node-fetch");

(async () => {
  let nextroute = "/";

  const basepath = "https://yourdomain.com";

  do {
    const req = await fetch(`${basepath}/${nextroute}`, {
      headers: {
        Accept: "application/json",
      },
    });

    const res = await req.json();

    console.log(res);

    if (res.next) {
      nextroute = res.next;
    } else {
      nextroute = null;
    }
  } while (nextroute !== null);
})();
