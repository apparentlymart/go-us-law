// Package bills contains a parser and object model for the XML schema used
// to publish legislation (i.e. bills) from the U.S. Congress.
//
// For more information on the format, see http://xml.house.gov/
//
// Bills in this format can either be obtained directly from the government
// system FDSys or via intermediaries with more convenient access interfaces,
// such as Govtrack: https://www.govtrack.us/developers/data
//
// The main entry points for this package are ParseBill and ParseBillBytes .
package bills
