<template>
  <div class="main-content">
    <div class="container">
      <form action="" @submit.prevent="handleProcess">
      <h3>{{process=='login'?'login':'register'}} form</h3>
        <div class="input">
          <Input
            v-focus
            :type="'text'"
            :placeholder="'Enter Email'"
            :id="'userEmail'"
          />
          <Input
            :type="'password'"
            :placeholder="'Enter Password'"
            :id="'password'"
          />
          <p class="note" :style="{color:messageColor}">{{ loginMessage }}</p>
        </div>
        <div class="button-group">
          <Button
            :text="process == 'login' ? 'Login' : 'Register'"
            :color="'white'"
            @click="handleProcess"
          />
          <button class="subButton" @click.prevent="redirect">
            {{process == 'login' ?'Register':'Login'}}
          </button>
        </div>
        <div class="decoration">
          <div class="holes">
            <div class="hole"></div>
            <div class="hole"></div>
            <div class="hole"></div>
            <div class="hole"></div>
            <div class="hole"></div>
          </div>
          <div class="hangers">
            <div class="hanger"></div>
            <div class="hanger"></div>
            <div class="hanger"></div>
            <div class="hanger"></div>
            <div class="hanger"></div>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { useStore } from "vuex";

import { ref } from "@vue/reactivity";

import Button from "../components/Button.vue";
import Input from "../components/Input.vue";
import axios from "axios";
export default {
  name: "Login",
  components: {
    Button,
    Input,
  },
  setup() {
    const store = useStore();
    const process = ref("login");
    const loginAction = (payload) => store.dispatch("act_login", payload);
    const loginMessage = ref("");
    const messageColor = ref("");
    const redirect = () => {
      if ((process.value == "login")) {
        process.value = "register";
      }else{
        process.value = "login";
      }
      loginMessage.value='';
    };
  

    const handleProcess = async () => {
      loginMessage.value = "";
      const formEntries = Array.prototype.slice.call(
        event.target.parentNode.parentNode.childNodes[1].childNodes
      );
      console.log(formEntries);
      const userEmail = formEntries.filter((e) => e.id == "userEmail")[0].value;
      const password = formEntries.filter((e) => e.id == "password")[0].value;
      let requestBody = { userEmail, userPassword: password };

      if ((process.value == "login")) {
        const response = await axios.post("/login-process", requestBody);
        if (typeof response.data == "string") {
          loginMessage.value = response.data;
          messageColor.value="red";
        } else {
          console.log(response.data);
          const payload = {
            email: response.data.userEmail,
            jwtToken: response.data.jwtToken,
          };
          loginAction(payload);
          window.location.href = "/";
        }
      } else {
        const response = await axios.post("/register-process", requestBody);
        if (response.data==''){
          loginMessage.value='Registered succesfully';
          messageColor.value='green';
        }
        else{
          loginMessage.value=response.data;
          messageColor.value='red';

        }
      }
    };
    return { process,loginMessage, messageColor, redirect, handleProcess };
  },
};
</script>
 
 <style lang="scss" scoped>
.main-content {
  background-image: url(https://media.architecturaldigest.com/photos/5f298951485b66fbbd9021cf/16:9/w_2560%2Cc_limit/IMG-22.jpg);
  background-color: $bg-color3;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  background-attachment: fixed;
  width: 100%;
  min-height: 100vh;
  overflow: auto;
  padding: 0;
  .container {
    //   margin: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    width: calc(100% - 450px);
    margin-left: 450px;
    min-height: 100vh;
    form {
      position: relative;
      border-radius: 50px;
      width: 500px;
      height: 500px;
      margin: auto;
      background: rgba(255, 255, 255, 0.6);
      padding: 1rem 3rem;
      display: flex;
      justify-content: space-evenly;
      align-items: center;
      flex-direction: column;
      h3{
        text-align: center;
        font-size: 1.5rem;
        margin: 2rem 0 0;
        width: 100%;
        text-transform: capitalize;
        text-decoration: underline;
      }
      .input {
        width: 100%;
        height: 250px;
        display: flex;
        justify-content: space-evenly;
        align-items: center;
        flex-direction: column;
      }
      .note {
        text-align: left;
        font-size: 1rem;
        width: 100%;
        height: 1.5rem;
        padding-left: 1rem;
        color: red;
        border-bottom: 1px solid gray;
      }
      .button-group {
        width: 100%;
      }

      button.subButton {
        width: 30%;
        display: block;
        background-color: none;
        color: black;
        font-weight: bold;
        text-decoration: underline;
        font-size: 1.1rem;
        font-style: italic;
        margin-left: auto;
        margin-top: 1rem;
      }
    }
    .decoration {
      position: absolute;
      top: 2%;
      left: 50%;
      transform: translate(-50%, 0);
      width: 80%;
      .holes,
      .hangers {
        display: flex;
        justify-content: space-around;
        width: 100%;
        .hole {
          --hole-size: 35px;
          width: var(--hole-size);
          height: var(--hole-size);
          background: none;
          border: 1px solid rgba(128, 128, 128, 0.3);
          border-radius: calc(var(--hole-size) / 2);
        }
      }
      .hangers {
        position: absolute;
        --hanger-height: 60px;
        --hanger-width: 25px;
        bottom: calc(50% - calc(var(--hanger-width) / 2));
        .hanger {
          width: var(--hanger-width);
          height: var(--hanger-height);
          border-radius: calc(var(--hanger-width) / 2);
          background: rgb(136, 99, 31);
        }
      }
    }
  }
}
</style>