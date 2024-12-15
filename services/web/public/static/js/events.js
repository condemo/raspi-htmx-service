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
    background: "#2DD4BF",
    color: "black",
    iconColor: "black",
    width: "15em",
    padding: ".9em .9em .9em"
  });
  Toast.fire({
    icon: "success",
    title: "Config Saved"
  });

})
