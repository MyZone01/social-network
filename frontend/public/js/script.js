// On page load or when changing themes, best to add inline in `head` to avoid FOUC
if (
  localStorage.theme === "dark" ||
  (!("theme" in localStorage) &&
    window.matchMedia("(prefers-color-scheme: dark)").matches)
) {
  document.documentElement.classList.add("dark");
} else {
  document.documentElement.classList.remove("dark");
}

// Whenever the user explicitly chooses light mode
localStorage.theme = "light";

// Whenever the user explicitly chooses dark mode
localStorage.theme = "dark";

// Whenever the user explicitly chooses to respect the OS preference
localStorage.removeItem("theme");

addEventListener("DOMContentLoaded", () => {
  console.log("CHANGING THEME");
  const toggleButton = document.getElementById("theme-mode");
  const body = document.body;

  console.log(body, toggleButton);

  toggleButton.addEventListener("click", () => {
    console.log("toogled")
    if (document.documentElement.classList.contains("dark")) {
      document.documentElement.classList.remove("dark")
    } else {
      document.documentElement.classList.add("dark")
    }
    // body.classList.toggle("dark");
  });
});


// }
// addEventListener('DOMContentLoaded', () => {
// // add post upload image
// document.getElementById('addPostUrl').addEventListener('change', function () {
//   if (this.files[0]) {
//     var picture = new FileReader();
//     picture.readAsDataURL(this.files[0]);
//     picture.addEventListener('load', function (event) {
//       document.getElementById('addPostImage').setAttribute('src', event.target.result);
//       document.getElementById('addPostImage').style.display = 'block';
//     });
//   }
// });

// // Create Status upload image
// document.getElementById('createStatusUrl').addEventListener('change', function () {
//   if (this.files[0]) {
//     var picture = new FileReader();
//     picture.readAsDataURL(this.files[0]);
//     picture.addEventListener('load', function (event) {
//       document.getElementById('createStatusImage').setAttribute('src', event.target.result);
//       document.getElementById('createStatusImage').style.display = 'block';
//     });
//   }
// });

// // create product upload image
// document.getElementById('createProductUrl').addEventListener('change', function () {
//   if (this.files[0]) {
//     var picture = new FileReader();
//     picture.readAsDataURL(this.files[0]);
//     picture.addEventListener('load', function (event) {
//       document.getElementById('createProductImage').setAttribute('src', event.target.result);
//       document.getElementById('createProductImage').style.display = 'block';
//     });
//   }
// });
// })
