<template>
    <page-card>
        <h2 class="head-register">Register Leave System</h2>
        <form action="">
            <div class="form-register">
                <label for="fullname">Fullname</label>
                <input type="text" name="fullname" id="fullname">
            </div>
            <div class="form-register">
                <label for="email">Email</label>
                <input type="emal" name="email" id="email">
            </div>
            <div class="form-register">
                <label for="password">Password</label>
                <input type="password" name="password" id="password">
            </div>
            <div class="form-register">
                <label for="password2">Confirm-Password</label>
                <input type="password" name="password2" id="password2">
            </div>
            <div class="form-register">
                <label for="password2">Confirm-Password</label>
                <input type="password" name="password2" id="password2">
            </div>
            <div>
                <page-button type="submit">
                    Register
                </page-button>
            </div>
        </form>
    </page-card>
</template>
<script>
export default {
    name: 'Register',
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

.form-register {
    margin: 1rem 0;
}

.head-register {
    background-image: linear-gradient(to right, #7abdff, #ff5f5f, #f364a4, #be5cff);
    -webkit-background-clip: text;
    color: transparent;
    background-size: 300%;
    background-position: -100%;
    animation: animatedText 2s infinite alternate-reverse;
}
@keyframes animatedText{
    to{
        background-position: 100%;
    }
}
</style>