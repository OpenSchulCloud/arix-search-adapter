/* This module contans the client for an ARIX interface.
 *
 */

package arix
 
 
/* This is the client interface for the arix api.
 *
 */
type ArixClient interface {
  
}


/* This is the interface to speak with an ARIX api via http or https.
 *
 */
type ArixHttpClient struct {
  
}


/* This endpoint uses an Arix Client to communicate and yields useful results.
 *
 */
type ArixApi struct {
  client ArixClient
}

