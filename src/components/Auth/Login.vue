<template>
    <page-card>
        <h2  class="head-login">Login Leave System</h2>
        <form action="">
            <div class="form-login">
                <label for="email">Email</label>
                <input type="email" name="email" id="email">
            </div>
            <div class="form-login">
                <label for="password">Password</label>
                <input type="password" name="password" id="password">
            </div>
            <div>
                <page-button type="submit">
                    Login
                </page-button>
            </div>
        </form>
    </page-card>
</template>
<script>
export default {
    name: 'Login',
}
</script>
<style scoped>
input {
    display: block;
    width: 100%;
    font: inherit;
    padding: 0.15rem;
    border: 1px solid #ccc;
}

input:focus {
    outline: none;
    border-color: #3a0061;
    background-color: #f7ebff;
}

.form-login {
    margin: 1rem 0;
}

.head-login {
    background-image: linear-gradient(to right, #00F260, #f79d00, #0575E6, #64f38c, #be5cff);
    -webkit-background-clip: text;
    color: transparent;
    background-size: 300%;
    background-position: -100%;
    animation: animatedText 5s infinite alternate-reverse;
}
@keyframes animatedText{
    to{
        background-position: 100%;
    }
}
</style>