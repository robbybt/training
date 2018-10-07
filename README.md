# Training Task
1. Make the training handler run as fast as possible
2. Add context cancel and timeout 3 second
3. Add hitcount on each request and add errorcount each error request happen
4. Add feature on background process which will check every 1 second check if errorcount > 10 then :
     - create a file for log the detail
     - clear hitcount & errorcount
     - pause this background process for 5 second
5. Add feature on background process for clear hitcount and errorcount every 10 second
6. Add fail over publish nsq retry 5 time with background process ( no need to wait the retry )

# Rule Task
1. Max 10 Go routine at one time ( The count is just only for go routine in func training handler )
2. Dont change constant response
