<script>
  import { CheckPasswords } from "../wailsjs/go/main/App.js";
  import {Greet} from '../wailsjs/go/main/App.js'

  let questionText = "Please enter your password to check below. Muliple entries allowed, separated by new line 👇"
  let name
  let pass_text

  function greet() {
    Greet(name).then(result => questionText = result)
  }
  function check_password() {
    CheckPasswords(pass_text).then(result => console.log(result))
  }
</script>

<main>
  <div class="result" id="result">{questionText}</div>
  <div class="input-box" id="input">
    <textarea bind:value={pass_text} class="input" id="pass_text" type="textarea"/>
    <button class="btn" on:click={check_password}>Check password</button>
  </div>
</main>

<style>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 100px;
    line-height: 15px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
