<script lang="ts">
import {onMount} from "svelte"
let result: String | null;
async function gatewayTest (){
        const resp = await fetch('http://localhost:8080',
        {method: "POST"}
        )
        const json = await resp.json()
        result = JSON.stringify(json, null, 2)
    }
async function searchTest (){
        const resp = await fetch('http://localhost:8080/searchtest'
        )
        const json = await resp.json()
        result = JSON.stringify(json, null, 2)
    }

</script>
<div>
    <h1>Personal Library Microservices Application</h1>
    <div class="action-container">
    <div class="button-container">
    <button class="menu-button" role="button" on:click={gatewayTest}>Gateway Test</button>
    <br>
    <br>
    <button class="menu-button" role="button" on:click={searchTest}>Search Test</button>
    <br>
    <br>
    <button class="menu-button" role="button">Library Test</button>
    </div>
    <div class="test-display">
    <pre>{result}</pre>
    </div>
    </div>

</div>

<style>
    .action-container{
            display: grid; 
            grid-template-columns: 1fr 1fr; 
            grid-template-rows: 1fr; 
            gap: 0px 5px; 
            grid-template-areas: ". ."; 
        }
    .button-container{
        display: block;
        border-style: solid;
        border-color: black;  
        border-radius: 1rem;
        padding: 20px;
    }
    .test-display{
        border-style: solid;
        border-color: black;  
        border-radius: 1rem;
        padding: 20px;
    }

    .menu-button {
        align-items: center;
        appearance: none;
        background-color: #FCFCFD;
        border-radius: 4px;
        border-width: 0;
        box-shadow: rgba(45, 35, 66, 0.4) 0 2px 4px,rgba(45, 35, 66, 0.3) 0 7px 13px -3px,#D6D6E7 0 -3px 0 inset;
        box-sizing: border-box;
        color: #36395A;
        cursor: pointer;
        display: inline-flex;
        font-family: "JetBrains Mono",monospace;
        height: 48px;
        justify-content: center;
        line-height: 1;
        list-style: none;
        overflow: hidden;
        padding-left: 16px;
        padding-right: 16px;
        position: relative;
        text-align: left;
        text-decoration: none;
        transition: box-shadow .15s,transform .15s;
        user-select: none;
        -webkit-user-select: none;
        touch-action: manipulation;
        white-space: nowrap;
        will-change: box-shadow,transform;
        font-size: 18px;
    }

    .menu-button:focus {
        box-shadow: #D6D6E7 0 0 0 1.5px inset, rgba(45, 35, 66, 0.4) 0 2px 4px, rgba(45, 35, 66, 0.3) 0 7px 13px -3px, #D6D6E7 0 -3px 0 inset;
    }

    .menu-button:hover {
        box-shadow: rgba(45, 35, 66, 0.4) 0 4px 8px, rgba(45, 35, 66, 0.3) 0 7px 13px -3px, #D6D6E7 0 -3px 0 inset;
        transform: translateY(-2px);
    }

    .menu-button:active {
        box-shadow: #D6D6E7 0 3px 7px inset;
        transform: translateY(2px);
    }
</style>
