package xyz.hoper

import org.jboss.logging.Logger
import java.util.*
import java.util.concurrent.atomic.AtomicInteger
import javax.ws.rs.GET
import javax.ws.rs.Path
import javax.ws.rs.Produces
import javax.ws.rs.ServerErrorException
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response

@Path("/logging-json")
class LoggingJsonResource {
    private val speed = AtomicInteger(0)
    private val random = Random()

    @GET
    @Path("faster")
    @Produces(MediaType.TEXT_PLAIN)
    fun faster(): String {
        val s = speed.addAndGet(random.nextInt(200))
        if (s > SPEED_OF_SOUND_IN_METER_PER_SECOND) {
            return "$s 💥 SONIC BOOOOOM!!!"
        }
        val message = "Your jet aircraft speed is $s m/s."
        LOG.info(message)
        return "$message Watch the logs..."
    }

    companion object {
        private val LOG = Logger.getLogger(LoggingJsonResource::class.java)
        private const val SPEED_OF_SOUND_IN_METER_PER_SECOND = 343
    }
}