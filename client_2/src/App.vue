<template>
  <div class="top-banner-overlay" :class="menuShowing ? 'is-moved' : ''">
    <h1>To home lovers</h1>
    <p>
      <cite>
        Home is where love resides, memories are created, friends and family
        belong, and laughter never ends.</cite
      >
    </p>
    <p class="author">- Unknown -</p>
  </div>
  <nav class="top-nav">
    <div class="menu-wrapper" :class="menuShowing ? 'is-opened' : ''">
      <ul class="menu">
        <li>
          <router-link to="/">home</router-link>
        </li>
        <li>
          <router-link to="/products">Products</router-link>
        </li>

        <li>
          <router-link to="/about">About us</router-link>
        </li>

        <li>
          <router-link to="/contact">contact us</router-link>
        </li>
      </ul>
      <button class="menu-close" aria-label="close menu" @click="closeMenu">
        ✕
      </button>
    </div>
    <div class="fixed-menu">
      <h2 class="logo">Home Decor</h2>
      <button class="menu-open" aria-label="open menu" @click="openMenu">
        &#x2630;
      </button>
      <ul class="socials">
        <li>
          <a href="">facebook</a>
        </li>
        <li>
          <a href="">twitter</a>
        </li>
        <li>
          <a href="">instagram</a>
        </li>
      </ul>
    </div>
  </nav>
  <div class="user-toolbar">
    <div class="item user-info" @mouseenter="showUserOptionList">
      <i class="far fa-user-circle"></i>
      <div></div>
      <ul class="hide" @mouseleave="hideUserOptionList">
        <div class="hanger"></div>
        <li v-if="userLogged">User profile</li>
        <li @click="userLogged?logout():login()" >{{userLogged?'Logout':'Login' }}</li>
      </ul>
    </div>
    <div class="item shopping-cart"><i class="fas fa-shopping-cart"></i></div>
  </div>
  <router-view />
</template>

<script>
import { ref } from "@vue/reactivity";
import { useStore } from "vuex";
import { computed } from '@vue/reactivity';

export default {
  name: "App",
  setup() {
    const store = useStore();
    const userLogged= computed(()=>store.getters.getStatus)
    const menuShowing = ref(false);
    const openMenu = () => {
      menuShowing.value = true;
    };
    const closeMenu = () => {
      menuShowing.value = false;
    };
    const showUserOptionList = (event) => {
      let childNodeArray = Array.prototype.slice.call(event.target.childNodes);
      childNodeArray = childNodeArray.filter((e) => {
        return e.nodeName.toLowerCase() != "text";
      });
      console.log(childNodeArray);
      const listElement = childNodeArray[childNodeArray.length - 1];
      listElement.classList.remove("hide");
    };
    const hideUserOptionList = (e) => {

      e.target.classList.add('hide');
      console.log(e.target.classList);
    };
    document.addEventListener("scroll", function () {
      const currentScrollCoord = window.scrollY;
      const scrollHistory = store.getters.getScrollY;
      const lastScrollCoord = scrollHistory[scrollHistory.length - 1];
      const windowHeight = window.innerHeight;
      let direction;
      console.log(lastScrollCoord, currentScrollCoord, windowHeight);
      if (currentScrollCoord >= lastScrollCoord) {
        direction = "down";
      } else {
        direction = "up";
      }
      console.log(direction);
      const upperLimit = Math.ceil(currentScrollCoord / windowHeight);
      const lowerLimit = Math.floor(currentScrollCoord / windowHeight);
      switch (direction) {
        case "down":
          if (
            currentScrollCoord > windowHeight * (upperLimit - 0.2) &&
            currentScrollCoord <= windowHeight * upperLimit
          ) {
            console.log("hit upper");
            window.scrollTo({
              left: 0,
              top: windowHeight * upperLimit,
              behavior: "smooth",
            });
          }
          break;
        case "up":
          if (
            currentScrollCoord < windowHeight * (lowerLimit + 0.2) &&
            currentScrollCoord >= windowHeight * lowerLimit
          ) {
            console.log("hit lower");
            window.scrollTo({
              left: 0,
              top: windowHeight * lowerLimit,
              behavior: "smooth",
            });
          }
      }
      store.dispatch("act_updateScrollY", currentScrollCoord);
    });
    const login= ()=>{
      console.log("hello");
      window.location.href="/login";
    }
    const logout= ()=>{
      const logoutAction = store.dispatch("act_logout");
      logoutAction();
    }
    return { userLogged,menuShowing, openMenu, closeMenu, showUserOptionList,hideUserOptionList,login ,logout};
  },

};
</script>

<style lang="scss">
* {
  box-sizing: border-box;
  // overflow: hidden;
}

body {
  padding: 0;
  margin: 0;
}

ul {
  list-style: none;
  padding-left: 0;
}

button {
  border: none;
  background: transparent;
  outline: none;
  cursor: pointer;
}
cite::before,
cite::after {
  content: '"';
}

button:active {
  color: black;
}

a {
  text-decoration: none;
  color: black;
}

body {
  font: normal 16px/1.5 Helvetica, sans-serif;
}

/* .top-banner
–––––––––––––––––––––––––––––––––––––––––––––––––– */

.top-banner-overlay {
  position: fixed;
  z-index: 90;
  left: 150px;
  padding: 1rem;
  color: white;
  top: 0;
  width: 300px;
  height: 100vh;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  flex-direction: column;
  transition: transform 0.7s;
}

.top-banner-overlay.is-moved {
  transform: translateX(350px);
}

.top-banner-overlay.is-moved::before {
  content: "";
  position: absolute;
  top: 0;
  bottom: 0;
  right: 100%;
  width: 20px;
  box-shadow: 0 0 10px black;
}

.top-nav li {
  padding: 0 0.5rem;
  margin-right: 2rem;
  text-transform: capitalize;
  &:hover {
    text-decoration: underline;
  }
  & + li {
    margin-top: 7px;
  }
}

/* .menu-wrapper
–––––––––––––––––––––––––––––––––––––––––––––––––– */
.top-nav .menu-wrapper {
  position: fixed;
  z-index: 90;
  top: 0;
  left: 0;
  bottom: 0;
  width: 350px;
  padding: 20px;
  transform: translateX(-200px);
  transition: transform 0.7s;
  background: $bg-color2;
}

.top-nav .menu-wrapper.is-opened {
  transform: translateX(150px);
}

.top-nav .menu-wrapper .menu {
  opacity: 0;
  transition: opacity 0.4s;
}

.top-nav .menu-wrapper.is-opened .menu {
  opacity: 1;
  transition-delay: 0.6s;
}

.top-nav .menu-wrapper .menu a {
  font-size: 1.2rem;
}

.top-nav .menu-wrapper .sub-menu {
  padding: 10px 0 0 7px;
}

.top-nav .menu-wrapper .menu-close {
  position: absolute;
  top: 20px;
  right: 20px;
  font-size: 1.6rem;
}

/* .fixed menu
–––––––––––––––––––––––––––––––––––––––––––––––––– */

.top-nav .fixed-menu {
  position: fixed;
  z-index: 100;

  top: 0;
  left: 0;
  bottom: 0;
  display: flex;
  flex-direction: column;
  width: 150px;
  padding: 20px;
  background: $bg-color1;
  h2 {
    font-size: 1.4rem;
  }
}

.top-nav .fixed-menu .menu-open {
  font-size: 1.8rem;
  text-align: left;
  margin: 30px 0 auto;
  width: 28px;
}

.user-toolbar {
  --height: 50px;
  --width: 200px;
  position: fixed;
  z-index: 110;
  top: 10px;
  right: 10px;
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: var(--height);
  width: var(--width);
  border-radius: calc(var(--height) / 2);
  background: rgba(255, 255, 255, 0.5);
  div.item {
    position: relative;
    --side: 40px;
    width: var(--side);
    height: var(--side);
    border-radius: calc(var(--side) / 2);
    // border: 1px solid gray;
    text-align: center;
    line-height: var(--side);
    font-size: 1.1rem;
    cursor: pointer;
    &:hover {
      background-color: rgba(255, 255, 255, 0.5);
      font-size: 1.2rem;
    }
  }
  div.user-info {
    ul {
      --background-color: rgba(255, 255, 255, 0.6);
      // display: none;
      position: absolute;
      right: 0;
      width: max-content;
      padding: 0.5rem;
      border-radius: 15px;
      font-size: 1rem;
      background: var(--background-color);
      transition: all 0.4s;
      div.hanger {
        position: absolute;
        bottom: 100%;
        left: 30%;
        background: var(--background-color);
        height: 26px;
        width: 80px;
        clip-path: polygon(10% 100%, 68% 0, 60% 100%);
      }
      li {
        padding: 0rem 0.8rem;
        text-align: left;
        border-radius: 5px;
        &:hover {
          background: white;
        }
      }
    }
    ul.hide {
      transform: scale(0.01);
      transform-origin: 80% -30%;
    }
  }
}
</style>