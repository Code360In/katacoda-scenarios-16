######################################################

######################################################
*** Settings ***
Library           SeleniumLibrary   run_on_failure=Nothing   timeout=30 seconds
Library           OperatingSystem
Library           Collections
Library           DateTime
Suite Teardown    Close all Connection Execute Teardown


*** Variables ***
${BROWSERNAME}         headlessfirefox
${URL}                 https://www.katacoda.com/
${T1}                  120

######################################################
*** Keywords ***


Close all Connection Execute Teardown
    Close Browser

Open Browser To Login Page
    [Arguments]     ${url} 
    Open Browser    ${url}    ${BROWSERNAME} 
    Maximize Browser Window
  


######################################################
*** Test Cases ***

Task Open Browse and Login
    Open Browser To Login Page   ${URL}

Task Connect to kotacoda scenario
    Open Browser To Login Page   https://www.katacoda.com/va/scenarios/

######################################################
