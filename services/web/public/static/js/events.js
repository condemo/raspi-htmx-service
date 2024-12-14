document.getElementById("config-section").addEventListener("htmx:afterRequest", function (evt) {

  const Toast = Swal.mixin({
    toast: true,
    position: "top-end",
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: false,
    didOpen: (toast) => {
      toast.onmouseenter = Swal.stopTimer;
      toast.onmouseleave = Swal.resumeTimer;
    },
    background: "green",
    color: "white"
  });
  Toast.fire({
    icon: "success",
    title: "Config Saved"
  });

})
