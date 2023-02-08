<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
     <div class="login-box">
        <h2>Leave System</h2>
        <form>
            <div class="user-box">
                <input type="text" name="" required="">
                <label for="">Username</label>
            </div>
            <div class="user-box">
                <input type="password" name="" required="">
                <label for="">Password</label>
            </div>
            <div class="button-form">
                <a id="submit" href="">Submit</a>
                <div id="register">
                    Don't have an account
                    <a href="">Register</a>
                </div>
            </div>
        </form>
    </div>
  <!-- <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div> -->
  
  <!-- <HelloWorld msg="Vite + Vue" /> -->
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Nunito+Sans:wght@200&display=swap');
/* :root {} */

* {
    box-sizing: border-box;
}

body {
    background: #131315;
    font-family: 'Nunito Sans', sans-serif;
}

p::selection {
    color: #000;
    background-color: #00657b;
}

.login-box {
    position: absolute;
    top: 15%; 
    Left: 37%;
    width: 400px;
    padding: 40px; 
    transform: translate(-50% -50%);
    background: #242826;
    box-shadow: 0 15px 25px rgba(143 124, 236, 0.7);
    border-radius: 10px;
}

.login-box h2 {
    margin: 30px;
    padding: 0;
    color: #fff;
    text-align: center;
}

.login-box .user-box input {
    position: relative;
    width: 100%;
    padding: 10px 0; 
    font-size: 16px;
    color: #fff;
    margin-bottom: 30px; 
    border: none;
    border-bottom: 1px solid #fff;
    outline: none; 
    background: transparent;
}

.login-box .user-box label {
    position: relative; 
    left: 0;
    top: -68px;
    padding: 10px 0;
    font-size: 16px; 
    color: #fff;
    pointer-events: none;
    transition: .5s;
}

.login-box .user-box input:focus~label,
.login-box .user-box input:valid~label {
    top: -85px;
    left: 0;
    color: #8F7CEC;
    font-size: 12px;
}

#submit {
    padding: 10px 20px;
    color: #CBBDDB; 
    text-transform: uppercase; 
    transition: .5s;
    border: 1px solid #8F7CEC;
    font-size: 16px;
    text-decoration: none;
    overflow: hidden;
    letter-spacing: 4px;
    margin: auto!;
}

#submit:hover {
    background: #8F7CEC;
    color: #fff;
    border-radius: 5px; 
    box-shadow: 0 0 5px #8F7CEC, 0 0 25px #8F7CEC, 0 0 50px #8F7CEC, 0 0 100px #8F7CEC;
}

.button-form {
    display: flex;
    flex-direction: row; 
    margin-top: 20px;
}

#register {
    font-size: 14px;
    text-decoration: none; 
    margin: auto; 
    text-align: center;
    color: #CBBDDB;
    width: 60%;
}

#register a {
    color: #8F7CEC;
    margin: auto;
    text-decoration: none;
}


    
/* .logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
} */
</style>
