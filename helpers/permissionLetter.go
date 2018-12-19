package helpers

import (
	"strings"

	"../model"
)

var xmlSource = (`
<document page-count="2">
<page number="1">
<table data-filename="CC102-converted-converted.pdf" data-page="1" data-table="1">
<tr><td colspan="4">OFFICE OF THE STUDENTS' WELFARE</td></tr>
<tr><td colspan="4">EVENT REQUEST FORM</td></tr>
<tr><td>CLUB/CHAPTER NAME:</td><td colspan="3">Clubname</td></tr>
<tr><td>EVENT NAME:</td><td colspan="2">Name</td><td></td></tr>
<tr><td>DATE:</td><td colspan="2">FromDate</td><td style="text-align: right">year of from</td></tr>
<tr><td>TIME:</td><td colspan="2">FromTime</td><td>ToTime</td></tr><tr><td colspan="4">DURATION: 2 Hours</td></tr>
<tr><td colspan="4">EVENT DESCRIPTION:</td></tr><tr><td colspan="4">Description</td></tr>
  
  <tr><td>EXPECTED PARTICIPANTS:</td><td style="text-align: right" colspan="3">ExpectedParticipants</td></tr>
  <tr><td>VENUE REQUIREMENT:</td><td colspan="2">Venue</td><td>Class Room: Yes.</td></tr>
  
  <tr><td>GUEST DETAILS (If Applicable):</td><td colspan="3">NA</td></tr>
  <tr><td>BUDGET (If Applicable):</td><td colspan="2">Budget</td><td style="text-align: right">-</td></tr>
  <tr><td>SPONSORS (If Applicable):</td><td colspan="2">NA</td><td style="text-align: right">-</td></tr> 
  <tr> <td colspan="2">STUDENT CO ORDINATOR DETAILS</td><td colspan="2">FACULTY CO ORDINATOR DETAILS</td></tr> 
  <tr><td>Name:</td><td>SCName</td><td>Name:</td><td>FCName</td></tr>
  <tr><td>Reg. No.</td><td>SCREgistrationNumber</td><td>Emp. ID.</td><td style="text-align: right">FCRegistrationNumber</td></tr>
  <tr><td>Mobile No.</td><td style="text-align: right">SCPhoneNumber</td><td>Mobile No.</td><td style="text-align: right">FCPhoneNumber</td>
  </tr><tr><td>Email:</td><td>SCEmail</td><td>Email:</td><td>FCEmail</td></tr>
  <tr><td>Signature:</td><td></td><td>Signature:</td><td></td></tr><tr><td></td><td></td><td>Seal:</td><td></td></tr>
  <tr><td>ADSW</td><td colspan="3">DSW</td></tr></table></page>


<page number="2">
<table data-filename="CC102-converted-converted.pdf" data-page="2" data-table="1">
<tr><td>Campus Engineer Office Request:<br/>CampusEngineerRequest</td></tr>
<tr><td style="text-align: right">1</td></tr><tr><td style="text-align: right">2</td></tr>
<tr><td style="text-align: right">3</td></tr><tr><td style="text-align: right">4</td></tr>
<tr><td style="text-align: right">5</td></tr><tr><td style="text-align: right">6</td></tr>
<tr><td>PRO Office Request:<br/>PROrequest</td></tr><tr><td style="text-align: right">1</td></tr>
<tr><td style="text-align: right">2</td></tr><tr><td style="text-align: right">3</td></tr>
<tr><td style="text-align: right">4</td></tr><tr><td style="text-align: right">5</td></tr>
<tr><td style="text-align: right">6</td></tr><tr><td>OTHER REQUESTS:</td></tr></table></page>
</document>

`)

func GenerateLetter(data model.Event) string {

	xmlSource = strings.Replace(xmlSource, "Clubname", data.ClubName, 1)
	xmlSource = strings.Replace(xmlSource, "Name", data.Name, 1)
	xmlSource = strings.Replace(xmlSource, "ToDate", data.ToDate, 1)
	xmlSource = strings.Replace(xmlSource, "FromDate", data.FromDate, 1)
	xmlSource = strings.Replace(xmlSource, "ToTime", data.ToTime, 1)
	xmlSource = strings.Replace(xmlSource, "FromTime", data.FromTime, 1)
	xmlSource = strings.Replace(xmlSource, "ExpectedParticipants", data.ExpectedParticipants, 1)
	xmlSource = strings.Replace(xmlSource, "Venue", data.Venue, 1)
	xmlSource = strings.Replace(xmlSource, "SCName", data.StudentCoordinator.Name, 1)
	xmlSource = strings.Replace(xmlSource, "FCName", data.FacultyCoordinator.Name, 1)
	xmlSource = strings.Replace(xmlSource, "SCPhoneNumber", data.StudentCoordinator.PhoneNumber, 1)
	xmlSource = strings.Replace(xmlSource, "FCPhoneNumber", data.FacultyCoordinator.PhoneNumber, 1)
	xmlSource = strings.Replace(xmlSource, "SCEmail", data.StudentCoordinator.Email, 1)
	xmlSource = strings.Replace(xmlSource, "FCEmail", data.FacultyCoordinator.Email, 1)
	xmlSource = strings.Replace(xmlSource, "SCRegistrationNumber", data.StudentCoordinator.RegistrationNumber, 1)
	xmlSource = strings.Replace(xmlSource, "FCRegistrationNumber", data.FacultyCoordinator.RegistrationNumber, 1)
	xmlSource = strings.Replace(xmlSource, "Description", data.Description, 1)

	return xmlSource

}
