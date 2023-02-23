<template>
    <div class="">
        <TheHeader :title="title"/>
        <div class="container">
            <Navigation />
            <Login />
            <Register />
        </div>
    </div>
    
    
</template>

<script>
import TheHeader from './components/layouts/TheHeader.vue';
import Navigation from './components/Auth/Navigation.vue';
import Login from './components/Auth/Login.vue';
import Register from './components/Auth/Register.vue';

export default {
    name : 'App',
    components:{
        Navigation,
        TheHeader,
        Login,
        Register,
    },
    data() {
        return {
            title: 'Vue Leave System'
        }
    },
}

</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Nunito+Sans:wght@200&display=swap');

* {
    margin: 0;
    box-sizing: border-box;
    font-family: 'Nunito Sans', sans-serif;
}

.container {
    width: 90%;
    margin: 0 auto;
}

/* body, html {} */
</style>
