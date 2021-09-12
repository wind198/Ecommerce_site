<template>
 <div class="main-content"> <button class="browse" @click="handleClick">
    Browse our shop<i class="fas fa-arrow-right"></i>
  </button></div>
</template>

<script>
export default {
  name: "Home",
  setup() {
    const handleClick = (event) => {
      let elem = event.target;
      console.log(elem.classList);
      elem.classList.add("flying");
      console.log(elem.classList);
setTimeout(() => {
    window.location.href="/products"
}, 1000);

    };
    return { handleClick };
  },
};
</script>

<style lang="scss" scoped>

.main-content {
  background-image: url(https://media.architecturaldigest.com/photos/5f298951485b66fbbd9021cf/16:9/w_2560%2Cc_limit/IMG-22.jpg);
  // background-color: $bg-color3;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  background-attachment: fixed;
  width: 100%;
  min-height: 100vh;
  overflow: auto;
  // margin-top: 100px;
}

button.browse {
  position: fixed;
  left: calc(50% + 225px);
  top: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  justify-content: center;
  align-items: center;
  background: white;
  opacity: 0.7;
  font-size: 2.2rem;
  padding: 1rem 1.2rem;
  i {
    font-size: 2.5rem;
    font-family: "Font Awesome 5 Free";
    color: rgb(237, 187, 80);
    font-weight: bold;
    margin-left: 0.5rem;
    pointer-events: none;
  }
  border-radius: 20px;
}
button.browse.flying {
  @keyframes flying {
    20% {
        opacity: 1;
        transform: translate(calc(-50% - 100px), -50%);
        
    }
    100% {
        transform: translate(calc(-50% + 1000px), -50%);
    }
  }
  animation: flying 1.5s;
}
</style>